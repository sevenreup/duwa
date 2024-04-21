package ast

import (
	"github.com/sevenreup/chewa/src/token"
)

type Postfix struct {
	Token    token.Token
	Operator string
}

func (pe *Postfix) TokenLiteral() string { return pe.Token.Literal }

// TODO: Print this one properly
func (pe *Postfix) String() string {
	return ""
}