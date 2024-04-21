package ast

import "github.com/sevenreup/chewa/src/token"

type Compound struct {
	Expression
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (fl *Compound) TokenLiteral() string { return fl.Token.Literal }

// TODO: Print this one properly
func (fl *Compound) String() string {
	return ""
}

