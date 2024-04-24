package ast

import (
	"testing"

	"github.com/sevenreup/chewa/src/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VariableDeclarationStatement{
				Type: token.Token{Type: token.INTEGER, Literal: "nambala"},
				Identifier: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "nambala myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
