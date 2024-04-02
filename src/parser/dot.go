package parser

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/token"
)

func (parser *Parser) dotExpression(left ast.Expression) ast.Expression {
	currentToken := parser.curToken
	currentPrecedence := parser.curPrecedence()

	parser.nextToken()

	if parser.peekTokenIs(token.OPENING_PAREN) {
		// Method
		expression := &ast.MethodExpression{Token: currentToken, Left: left}
		expression.Method = parser.parseExpression(currentPrecedence)

		parser.nextToken()

		expression.Arguments = parser.parseExpressionList(token.CLOSING_PAREN)

		return expression
	}

	// TODO: Add logic for handling properties "class.me"

	return nil
}
