package ast

import "github.com/sevenreup/duwa/src/token"

type ImportExpression struct {
	Expression
	Token token.Token
	Path  *StringLiteral
}
