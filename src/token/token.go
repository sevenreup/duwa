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
	EQUAL_TO                 = "=="
	NOT_EQUAL_TO             = "!="

	ILLEGAL

	COMMENT           = "//"
	MULTILINE_COMMENT = "/*"

	IDENT = "identifier"

	// Keywords
	INTEGER  = "INTEGER"
	STRING   = "STRING"
	FUNCTION = "function"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
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

var keywords = map[string]TokenType{
	"nambala": INTEGER,
	"mawu":    STRING,
	"zoona":   TRUE,
	"bodza":   FALSE,
	"ngati":   IF,
	"kapena":  ELSE,
	"bweza":   RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
