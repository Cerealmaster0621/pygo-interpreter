package tokenizer

import (
	"bufio"
	"io"
	"unicode/utf8"
)

// TODO - Strings, F-Strings, Triple-Quoted Strings, Kommentare, Whitespaces, Indentation

// type Tokenizer -
type Tokenizer struct {
	reader    *bufio.Reader
	lineCount int
}

// NewTokenizer - Wrapper for NewReader
func NewTokenizer(r io.Reader) *Tokenizer {
	return &Tokenizer{
		reader:    bufio.NewReader(r),
		lineCount: 0,
	}
}

// ReadLine reads line by line until end of the lines
func (t *Tokenizer) ReadLine() ([]byte, error) {
	line, err := t.reader.ReadBytes('\n')
	if err == io.EOF && len(line) > 0 {
		// Letzte Zeile ohne Newline
		return line, nil
	}
	return line, err
}

// Tokenize liest die gesamte Eingabe ein und zerlegt sie in TokenInfo.
func (t *Tokenizer) Tokenize() ([]TokenInfo, error) {
	var tokens []TokenInfo
	lineNumber := 0

	for {
		line, err := t.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return tokens, err
		}
		lineNumber++
		tokens = append(tokens, t.tokenizeLine(line, lineNumber)...)
	}

	// Abschließender ENDMARKER
	tokens = append(tokens, TokenInfo{
		Type:   ENDMARKER,
		String: "",
		Start:  Position{Row: lineNumber + 1, Col: 0},
		End:    Position{Row: lineNumber + 1, Col: 0},
		Line:   "",
	})

	return tokens, nil
}

// tokenizeLine zerlegt eine einzelne Zeile in Tokens.
func (t *Tokenizer) tokenizeLine(line []byte, lineNumber int) []TokenInfo {
	var tokens []TokenInfo
	original := string(line)
	col := 0

	input := line
	for len(input) > 0 {
		// Entferne Leerzeichen am Anfang
		if loc := whitespaceRegex.FindIndex(input); loc != nil && loc[0] == 0 {
			spaces := loc[1]
			input = input[spaces:]
			col += spaces
			continue
		}

		// NAME
		if loc := nameRegex.FindIndex(input); loc != nil && loc[0] == 0 {
			val := input[0:loc[1]]
			tok := TokenInfo{
				Type:   NAME,
				String: string(val),
				Start:  Position{Row: lineNumber, Col: col},
				End:    Position{Row: lineNumber, Col: col + len(val)},
				Line:   original,
			}
			tokens = append(tokens, tok)
			col += len(val)
			input = input[len(val):]
			continue
		}

		// NUMBER
		if loc := IntnumberRegex.FindIndex(input); loc != nil && loc[0] == 0 {
			val := input[0:loc[1]]
			tok := TokenInfo{
				Type:   NUMBER,
				String: string(val),
				Start:  Position{Row: lineNumber, Col: col},
				End:    Position{Row: lineNumber, Col: col + len(val)},
				Line:   original,
			}
			tokens = append(tokens, tok)
			col += len(val)
			input = input[len(val):]
			continue
		}

		// OPERATOR
		if loc := opRegex.FindIndex(input); loc != nil && loc[0] == 0 {
			val := input[0:loc[1]]
			tokType := OP
			// Versuche exakten Operator-Typ aus ExactTokenTypes zu bestimmen:
			if exact, ok := ExactTokenTypes[string(val)]; ok {
				tokType = exact
			}

			tok := TokenInfo{
				Type:   tokType,
				String: string(val),
				Start:  Position{Row: lineNumber, Col: col},
				End:    Position{Row: lineNumber, Col: col + len(val)},
				Line:   original,
			}
			tokens = append(tokens, tok)
			col += len(val)
			input = input[len(val):]
			continue
		}

		// STRING oder andere Fälle:
		// Hier könntest du weitere Regex einsetzen, um Strings oder Kommentare zu erkennen.

		// Wenn kein Pattern passt, konsumieren wir ein einzelnes Zeichen als ERRORTOKEN
		r, size := utf8.DecodeRune(input)
		if r == utf8.RuneError && size == 1 {
			// Ungültiges UTF-8
			val := input[0:1]
			tok := TokenInfo{
				Type:   ERRORTOKEN,
				String: string(val),
				Start:  Position{Row: lineNumber, Col: col},
				End:    Position{Row: lineNumber, Col: col + 1},
				Line:   original,
			}
			tokens = append(tokens, tok)
			col += 1
			input = input[1:]
		} else {
			// Konsumiere ein Zeichen als Operator oder anderen Token
			val := input[0:size]
			tok := TokenInfo{
				Type:   ERRORTOKEN,
				String: string(val),
				Start:  Position{Row: lineNumber, Col: col},
				End:    Position{Row: lineNumber, Col: col + size},
				Line:   original,
			}
			tokens = append(tokens, tok)
			col += size
			input = input[size:]
		}
	}

	// Jede Zeile endet mit einem NEWLINE-Token, falls benötigt.
	if len(line) > 0 && line[len(line)-1] == '\n' {
		tokens = append(tokens, TokenInfo{
			Type:   NEWLINE,
			String: "\n",
			Start:  Position{Row: lineNumber, Col: col},
			End:    Position{Row: lineNumber, Col: col + 1},
			Line:   original,
		})
	}

	return tokens
}
