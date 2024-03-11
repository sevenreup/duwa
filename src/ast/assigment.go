package ast

import (
	"bytes"
	"github.com/sevenreup/chewa/src/token"
)

type AssigmentStatement struct {
	Token token.Token // the token.Nambala token
	Name  *Identifier
	Value Expression
}

func (ls *AssigmentStatement) statementNode()       {}
func (ls *AssigmentStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *AssigmentStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.Value)
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
