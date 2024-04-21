package ast

import (
	"bytes"
)

type AssigmentStatement struct {
	Statement
	Identifier *Identifier
	Value      Expression
}

func (ls *AssigmentStatement) TokenLiteral() string { return ls.Identifier.TokenLiteral() }

func (ls *AssigmentStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Identifier.Value)
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
