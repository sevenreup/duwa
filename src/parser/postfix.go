package parser

import (
	"github.com/sevenreup/duwa/src/ast"
)

func (parser *Parser) parsePostfixExpression() ast.Expression {
	return &ast.PostfixExpression{
		Token:    parser.previousToken,
		Operator: parser.curToken.Literal,
	}
}
