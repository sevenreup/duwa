package parser

import (
	"github.com/sevenreup/duwa/src/ast"
	"github.com/sevenreup/duwa/src/token"
)

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}
	if !p.expectPeek(token.OPENING_PAREN) {
		return nil
	}
	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)
	if !p.expectPeek(token.CLOSING_PAREN) {
		return nil
	}
	if !p.expectPeek(token.OPENING_BRACE) {
		return nil
	}
	expression.Consequence = p.parseBlockStatement()
	if p.peekTokenIs(token.ELSE) {
		p.nextToken()
		if !p.expectPeek(token.OPENING_BRACE) {
			return nil
		}
		expression.Alternative = p.parseBlockStatement()
	}
	return expression
}
