package ast

import "github.com/sevenreup/chewa/src/token"

type ForExpression struct {
	Token       token.Token
	Identifier  *Identifier
	Initializer Statement
	Condition   Expression
	Increment   Statement
	Block       *BlockStatement
}

func (ie *ForExpression) expressionNode() {}

func (ie *ForExpression) TokenLiteral() string { return ie.Token.Literal }

// TODO: Print this one properly
func (ie *ForExpression) String() string {
	return ""
}
