package parser

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/token"
)

func (p *Parser) parseAssignmentStatement() *ast.AssigmentStatement {
	stmt := &ast.AssigmentStatement{
		Identifier: &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal},
	}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
