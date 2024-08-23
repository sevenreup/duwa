package ast

import (
	"bytes"
	"github.com/sevenreup/duwa/src/token"
)

type IndexExpression struct {
	Expression
	AssignmentNode
	Token token.Token // The [ token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}
