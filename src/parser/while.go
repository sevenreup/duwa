package parser

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/token"
)

func (parser *Parser) parseWhileExpression() ast.Expression {
	expression := &ast.WhileExpression{Token: parser.curToken}

	if !parser.expectPeek(token.OPENING_PAREN) {
		return nil
	}

	parser.nextToken()
	expression.Condition = parser.parseExpression(LOWEST)

	if !parser.expectPeek(token.CLOSING_PAREN) {
		return nil
	}

	if !parser.expectPeek(token.OPENING_BRACE) {
		return nil
	}

	expression.Consequence = parser.parseBlockStatement()

	return expression
}
