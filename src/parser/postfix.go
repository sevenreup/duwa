package parser

import (
	"github.com/sevenreup/chewa/src/ast"
)

func (parser *Parser) parsePostfixExpression() ast.Expression {
	return &ast.PostfixExpression{
		Token:    parser.previousToken,
		Operator: parser.curToken.Literal,
	}
}
