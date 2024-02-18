package scanner

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

type TokenType int

const (
	// Single character tokens
	TOKEN_EOF     TokenType = iota
	TOKEN_MINUS             // -
	TOKEN_STAR              // *
	TOKEN_DIVIDE            // /
	TOKEN_PLUS              // +
	SEMI                    // ;
	GREATER_THAN            // >
	LESS_THAN               // <
	ASSIGN                  // =
	COLON                   // :
	COMMA                   // ,
	OPENING_BRACE           // {
	CLOSING_BRACE           // }
	OPENING_PAREN           // (
	CLOSING_PAREN           // )
	FULL_STOP               // .

	// One or two character token
	GREATER_THAN_OR_EQUAL_TO // >=
	LESS_THAN_OR_EQUAL_TO    // <=
	NOT_EQUAL_TO             // !=

	TOKEN_ILLEGAL
	INT
	STRING

	COMMENT           // // (comment)
	MULTILINE_COMMENT // /* (comment)

	IDENT

	QUERY_VARIALBE // $ (query variable)
)

type Position struct {
	Line   int
	Column int
}

type Scanner struct {
	r    *bufio.Reader
	pos  Position
	file *os.File
}

type TokenInfo struct {
	Position Position
	Token    TokenType
	Literal  string
}

func NewScanner(path string) *Scanner {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return &Scanner{
		r:    bufio.NewReader(file),
		pos:  Position{Line: 1, Column: 0},
		file: file,
	}
}

func (l *Scanner) Close() {
	l.file.Close()
}

func (l *Scanner) ReadTokens() (Position, TokenType, string) {
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, TOKEN_EOF, ""
			}
			panic(err)
		}

		l.pos.Column++
		switch r {
		case '\n':
			l.resetPosition()
		case ';':
			return l.pos, SEMI, ";"
		case '"', '\'':
			return l.ReadString()
		case '+':
			return l.pos, TOKEN_PLUS, "+"
		case '-':
			return l.pos, TOKEN_MINUS, "-"
		case ':':
			return l.pos, COLON, ":"
		case '/':
			nextRune := l.Peek()
			if nextRune == '/' {
				l.Next()
				return l.ReadComment()
			} else if nextRune == '*' {
				var builder strings.Builder
				builder.WriteString("/")
				for {
					r, _, err := l.r.ReadRune()
					if err != nil {
						if err == io.EOF {
							return l.pos, TOKEN_EOF, ""
						}
						panic(err)
					}
					l.pos.Column++
					if r == '*' {
						nextRune = l.Peek()
						if nextRune == '/' {
							l.Next()
							break
						}
					}
					builder.WriteString(string(r))
				}
				builder.WriteString("*/")
				return l.pos, MULTILINE_COMMENT, builder.String()
			}
			return l.pos, TOKEN_DIVIDE, "/"
		case '{':
			return l.pos, OPENING_BRACE, "{"
		case '}':
			return l.pos, CLOSING_BRACE, "}"
		case '(':
			return l.pos, OPENING_PAREN, "("
		case ')':
			return l.pos, CLOSING_PAREN, ")"
		case '.':
			return l.pos, FULL_STOP, "."
		case ',':
			return l.pos, COMMA, ","
		case '=':
			return l.pos, ASSIGN, "="
		case '>':
			nextRune := l.Peek()
			if nextRune == '=' {
				l.Next()
				return l.pos, GREATER_THAN_OR_EQUAL_TO, ">="
			}
			return l.pos, GREATER_THAN, ">"
		case '<':
			nextRune := l.Peek()
			if nextRune == '=' {
				l.Next()
				return l.pos, LESS_THAN_OR_EQUAL_TO, "<="
			}
			return l.pos, LESS_THAN, "<"
		case '!':
			nextRune := l.Peek()
			if nextRune == '=' {
				l.Next()
				return l.pos, NOT_EQUAL_TO, "!="
			}
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsLetter(r) {
				return l.ReadIdentifier(r)
			} else if unicode.IsDigit(r) {
				return l.ReadNumber(r)
			}
			return l.pos, STRING, string(r)
		}
	}
}

func (l *Scanner) Next() rune {
	r, _, err := l.r.ReadRune()
	if err != nil {
		return rune(0)
	}
	l.pos.Column++
	return r
}

func (l *Scanner) Peek() rune {
	r, _, _ := l.r.ReadRune()
	l.r.UnreadRune()
	return r
}

func (l *Scanner) ReadComment() (Position, TokenType, string) {
	rawString := ""
	var newPos Position
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, TOKEN_EOF, ""
			}
			panic(err)
		}
		l.pos.Column++
		if r == '\n' {
			newPos = l.pos
			l.resetPosition()
			break
		} else {
			rawString += string(r)
		}
	}
	return newPos, COMMENT, rawString
}

func (l *Scanner) ReadString() (Position, TokenType, string) {
	rawString := ""
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, TOKEN_EOF, ""
			}
			panic(err)
		}
		l.pos.Column++
		if r == '"' || r == '\'' {
			break
		} else {
			rawString += string(r)
		}
	}
	return l.pos, STRING, rawString
}

func (l *Scanner) ReadNumber(current rune) (Position, TokenType, string) {
	number := string(current)
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, TOKEN_EOF, ""
			}
			panic(err)
		}
		l.pos.Column++
		if unicode.IsDigit(r) {
			number += string(r)
		} else {
			l.r.UnreadRune()
			break
		}
	}
	return l.pos, INT, number
}

func (l *Scanner) ReadIdentifier(current rune) (Position, TokenType, string) {
	identifier := string(current)
	for {
		r, _, err := l.r.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, TOKEN_EOF, ""
			}
			panic(err)
		}
		l.pos.Column++
		if validIdentifierSymbol(r) {
			identifier += string(r)
		} else {
			l.r.UnreadRune()
			break
		}
	}
	return l.pos, IDENT, identifier
}

func validIdentifierSymbol(symbol rune) bool {
	return unicode.IsLetter(symbol) || unicode.IsDigit(symbol) || symbol == '_'
}

func (l *Scanner) resetPosition() {
	l.pos.Line++
	l.pos.Column = 0
}

func (l *Scanner) AccumTokens() []TokenInfo {
	var tokens []TokenInfo
	for {
		pos, tok, lit := l.ReadTokens()
		if tok == TOKEN_EOF {
			break
		}
		info := TokenInfo{
			Position: pos,
			Token:    tok,
			Literal:  lit,
		}
		tokens = append(tokens, info)
	}

	return tokens
}
