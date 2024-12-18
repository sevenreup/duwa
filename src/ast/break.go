package ast

import "github.com/sevenreup/duwa/src/token"

type BreakStatement struct {
	Statement
	Token token.Token
}

func (node *BreakStatement) TokenLiteral() string { return node.Token.Literal }
func (node *BreakStatement) String() string {
	return "siya"
}
