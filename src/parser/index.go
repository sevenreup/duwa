package parser

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/token"
)

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Token: p.curToken, Left: left}
	p.nextToken()
	exp.Index = p.parseExpression(LOWEST)
	if !p.expectPeek(token.CLOSING_BRACKET) {
		return nil
	}
	return exp
}
