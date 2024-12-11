package tokenizer

// TokenType represents token type as string
type TokenType string

const (
	ENDMARKER        TokenType = "ENDMARKER"
	NAME             TokenType = "NAME"
	NUMBER           TokenType = "NUMBER"
	STRING           TokenType = "STRING"
	NEWLINE          TokenType = "NEWLINE"
	INDENT           TokenType = "INDENT"
	DEDENT           TokenType = "DEDENT"
	LPAR             TokenType = "LPAR"
	RPAR             TokenType = "RPAR"
	LSQB             TokenType = "LSQB"
	RSQB             TokenType = "RSQB"
	COLON            TokenType = "COLON"
	COMMA            TokenType = "COMMA"
	SEMI             TokenType = "SEMI"
	PLUS             TokenType = "PLUS"
	MINUS            TokenType = "MINUS"
	STAR             TokenType = "STAR"
	SLASH            TokenType = "SLASH"
	VBAR             TokenType = "VBAR"
	AMPER            TokenType = "AMPER"
	LESS             TokenType = "LESS"
	GREATER          TokenType = "GREATER"
	EQUAL            TokenType = "EQUAL"
	DOT              TokenType = "DOT"
	PERCENT          TokenType = "PERCENT"
	LBRACE           TokenType = "LBRACE"
	RBRACE           TokenType = "RBRACE"
	EQEQUAL          TokenType = "EQEQUAL"
	NOTEQUAL         TokenType = "NOTEQUAL"
	LESSEQUAL        TokenType = "LESSEQUAL"
	GREATEREQUAL     TokenType = "GREATEREQUAL"
	TILDE            TokenType = "TILDE"
	CIRCUMFLEX       TokenType = "CIRCUMFLEX"
	LEFTSHIFT        TokenType = "LEFTSHIFT"
	RIGHTSHIFT       TokenType = "RIGHTSHIFT"
	DOUBLESTAR       TokenType = "DOUBLESTAR"
	PLUSEQUAL        TokenType = "PLUSEQUAL"
	MINEQUAL         TokenType = "MINEQUAL"
	STAREQUAL        TokenType = "STAREQUAL"
	SLASHEQUAL       TokenType = "SLASHEQUAL"
	PERCENTEQUAL     TokenType = "PERCENTEQUAL"
	AMPEREQUAL       TokenType = "AMPEREQUAL"
	VBAREQUAL        TokenType = "VBAREQUAL"
	CIRCUMFLEXEQUAL  TokenType = "CIRCUMFLEXEQUAL"
	LEFTSHIFTEQUAL   TokenType = "LEFTSHIFTEQUAL"
	RIGHTSHIFTEQUAL  TokenType = "RIGHTSHIFTEQUAL"
	DOUBLESTAREQUAL  TokenType = "DOUBLESTAREQUAL"
	DOUBLESLASH      TokenType = "DOUBLESLASH"
	DOUBLESLASHEQUAL TokenType = "DOUBLESLASHEQUAL"
	AT               TokenType = "AT"
	ATEQUAL          TokenType = "ATEQUAL"
	RARROW           TokenType = "RARROW"
	ELLIPSIS         TokenType = "ELLIPSIS"
	COLONEQUAL       TokenType = "COLONEQUAL"
	EXCLAMATION      TokenType = "EXCLAMATION"
	OP               TokenType = "OP"
	TYPE_IGNORE      TokenType = "TYPE_IGNORE"
	TYPE_COMMENT     TokenType = "TYPE_COMMENT"
	SOFT_KEYWORD     TokenType = "SOFT_KEYWORD"
	FSTRING_START    TokenType = "FSTRING_START"
	FSTRING_MIDDLE   TokenType = "FSTRING_MIDDLE"
	FSTRING_END      TokenType = "FSTRING_END"
	COMMENT          TokenType = "COMMENT"
	NL               TokenType = "NL"
	ERRORTOKEN       TokenType = "ERRORTOKEN"
	ENCODING         TokenType = "ENCODING"
)

// Position represents position of token
type Position struct {
	Row int
	Col int
}

// TokenInfo includes entire information about token
type TokenInfo struct {
	Type   TokenType
	String string
	Start  Position
	End    Position
	Line   string
}

// TokenName map TokenType to readable name
var TokenName = map[TokenType]string{
	ENDMARKER:        "End of Input",
	NAME:             "Name",
	NUMBER:           "Number",
	STRING:           "String",
	NEWLINE:          "Newline",
	INDENT:           "Indentation",
	DEDENT:           "Dedentation",
	LPAR:             "Left Parenthesis",
	RPAR:             "Right Parenthesis",
	LSQB:             "Left Square Bracket",
	RSQB:             "Right Square Bracket",
	COLON:            "Colon",
	COMMA:            "Comma",
	SEMI:             "Semicolon",
	PLUS:             "Plus",
	MINUS:            "Minus",
	STAR:             "Star",
	SLASH:            "Slash",
	VBAR:             "Vertical Bar",
	AMPER:            "Ampersand",
	LESS:             "Less",
	GREATER:          "Greater",
	EQUAL:            "Equal",
	DOT:              "Dot",
	PERCENT:          "Percent",
	LBRACE:           "Left Brace",
	RBRACE:           "Right Brace",
	EQEQUAL:          "Equal Equal",
	NOTEQUAL:         "Not Equal",
	LESSEQUAL:        "Less Equal",
	GREATEREQUAL:     "Greater Equal",
	TILDE:            "Tilde",
	CIRCUMFLEX:       "Circumflex",
	LEFTSHIFT:        "Left Shift",
	RIGHTSHIFT:       "Right Shift",
	DOUBLESTAR:       "Double Star",
	PLUSEQUAL:        "Plus Equal",
	MINEQUAL:         "Minus Equal",
	STAREQUAL:        "Star Equal",
	SLASHEQUAL:       "Slash Equal",
	PERCENTEQUAL:     "Percent Equal",
	AMPEREQUAL:       "Ampersand Equal",
	VBAREQUAL:        "VBar Equal",
	CIRCUMFLEXEQUAL:  "Circumflex Equal",
	LEFTSHIFTEQUAL:   "Left Shift Equal",
	RIGHTSHIFTEQUAL:  "Right Shift Equal",
	DOUBLESTAREQUAL:  "Double Star Equal",
	DOUBLESLASH:      "Double Slash",
	DOUBLESLASHEQUAL: "Double Slash Equal",
	AT:               "At Symbol",
	ATEQUAL:          "At Equal",
	RARROW:           "Right Arrow",
	ELLIPSIS:         "Ellipsis",
	COLONEQUAL:       "Colon Equal",
	EXCLAMATION:      "Exclamation",
	OP:               "Operator",
	TYPE_IGNORE:      "Type Ignore",
	TYPE_COMMENT:     "Type Comment",
	SOFT_KEYWORD:     "Soft Keyword",
	FSTRING_START:    "F-String Start",
	FSTRING_MIDDLE:   "F-String Middle",
	FSTRING_END:      "F-String End",
	COMMENT:          "Comment",
	NL:               "Non-Logical Newline",
	ERRORTOKEN:       "Error Token",
	ENCODING:         "Encoding",
}

// ExactTokenTypes map each operattors/symbole with exact token type
var ExactTokenTypes = map[string]TokenType{
	"!":   EXCLAMATION,
	"!=":  NOTEQUAL,
	"%":   PERCENT,
	"%=":  PERCENTEQUAL,
	"&":   AMPER,
	"&=":  AMPEREQUAL,
	"(":   LPAR,
	")":   RPAR,
	"*":   STAR,
	"**":  DOUBLESTAR,
	"**=": DOUBLESTAREQUAL,
	"*=":  STAREQUAL,
	"+":   PLUS,
	"+=":  PLUSEQUAL,
	",":   COMMA,
	"-":   MINUS,
	"-=":  MINEQUAL,
	"->":  RARROW,
	".":   DOT,
	"...": ELLIPSIS,
	"/":   SLASH,
	"//":  DOUBLESLASH,
	"//=": DOUBLESLASHEQUAL,
	"/=":  SLASHEQUAL,
	":":   COLON,
	":=":  COLONEQUAL,
	";":   SEMI,
	"<":   LESS,
	"<<":  LEFTSHIFT,
	"<<=": LEFTSHIFTEQUAL,
	"<=":  LESSEQUAL,
	"=":   EQUAL,
	"==":  EQEQUAL,
	">":   GREATER,
	">=":  GREATEREQUAL,
	">>":  RIGHTSHIFT,
	">>=": RIGHTSHIFTEQUAL,
	"@":   AT,
	"@=":  ATEQUAL,
	"[":   LSQB,
	"]":   RSQB,
	"^":   CIRCUMFLEX,
	"^=":  CIRCUMFLEXEQUAL,
	"{":   LBRACE,
	"|":   VBAR,
	"|=":  VBAREQUAL,
	"}":   RBRACE,
	"~":   TILDE,
}

const NT_OFFSET = 256

// Help functions
func IsEOF(tokenType TokenType) bool {
	return tokenType == ENDMARKER
}

func IsTerminal(tokenType int) bool {
	return tokenType < 256
}
