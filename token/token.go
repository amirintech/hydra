package token

// token types
type TokeType string

const (
	// syntax
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers & literals
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	PLUS         = "+"
	MINUS        = "-"
	ASTERISK     = "*"
	SLASH        = "/"
	BANG         = "!"
	LESS_THAN    = "<"
	GREATER_THAN = ">"
	ASSIGN       = "="
	EQUAL        = "=="
	NOT_EQUAL    = "!="

	// delimiters
	COMA      = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

// token struct
type Token struct {
	Type    TokeType
	Literal string
}

// keywords map
var keywords = map[string]TokeType{
	"let":    LET,
	"fn":     FUNCTION,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdentifier(ident string) TokeType {
	if keyword, ok := keywords[ident]; ok {
		return keyword
	}

	return IDENT
}
