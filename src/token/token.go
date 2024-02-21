package token

type TokenType string

const (
	// Single character tokens
	EOF           = "EOF"
	MINUS         = "-"
	STAR          = "*"
	DIVIDE        = "/"
	PLUS          = "+"
	SEMICOLON     = ";"
	GREATER_THAN  = ">"
	LESS_THAN     = "<"
	ASSIGN        = "="
	COLON         = ":"
	COMMA         = ","
	OPENING_BRACE = "{"
	CLOSING_BRACE = "}"
	OPENING_PAREN = "("
	CLOSING_PAREN = ")"
	FULL_STOP     = "."

	// One or two character token
	GREATER_THAN_OR_EQUAL_TO = ">="
	LESS_THAN_OR_EQUAL_TO    = "<="
	NOT_EQUAL_TO             = "!="

	ILLEGAL
	INT
	STRING

	COMMENT           = "//"
	MULTILINE_COMMENT = "/*"

	IDENT = "identifier"

	// Keywords
	FUNCTION = "function"
	INTEGER  = "interger"
)

type Position struct {
	Line   int
	Column int
}

type Token struct {
	Type    TokenType
	Literal string
	Pos     Position
}
