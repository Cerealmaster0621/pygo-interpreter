package tokenizer

type TokenInfo struct {
	Type   TokenType
	String string
	Start  Position
	End    Position
	Line   string
}

// type TokenType

type TokenType string

const (
	ENDMARKER  TokenType = "ENDMARKER"
	NAME       TokenType = "NAME"
	NUMBER     TokenType = "NUMBER"
	STRING     TokenType = "STRING"
	NEWLINE    TokenType = "NEWLINE"
	INDENT     TokenType = "INDENT"
	DEDENT     TokenType = "DEDENT"
	OPERATOR   TokenType = "OPERATOR"
	COMMENT    TokenType = "COMMENT"
	NL         TokenType = "NL"
	ERRORTOKEN TokenType = "ERRORTOKEN"
)

// type Position

type Position struct {
	Row int
	Col int
}

// type TokenName map TokenType as readable name

var TokenName = map[TokenType]string{
	ENDMARKER:  "End of Input",
	NAME:       "Name",
	NUMBER:     "Number",
	STRING:     "String",
	NEWLINE:    "Newline",
	INDENT:     "Indentation",
	DEDENT:     "Dedentation",
	OPERATOR:   "Operator",
	COMMENT:    "Comment",
	NL:         "Non-Logical Newline",
	ERRORTOKEN: "Error Token",
}

// type ExactTokenTypes map operations with token type

var ExactTokenTypes = map[string]TokenType{
	"!":  "EXCLAMATION",
	"!=": "NOTEQUAL",
	"%":  "PERCENT",
	"+":  "PLUS",
	"=":  "EQUAL",
}

// Helpfunctions for Token classification

func IsEOF(tokenType TokenType) bool {
	return tokenType == ENDMARKER
}

func IsTerminal(tokenType int) bool {
	return tokenType < 256
}
