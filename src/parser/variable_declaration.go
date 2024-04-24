package parser

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/token"
)

func (p *Parser) parseVariableDeclarationStatement() *ast.VariableDeclarationStatement {
	stmt := &ast.VariableDeclarationStatement{Type: p.curToken}

	if p.peekTokenIs(token.OPENING_BRACKET) {
		p.nextToken()
		if !p.peekTokenIs(token.CLOSING_BRACKET) {
			return nil
		}
		p.nextToken()
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Identifier = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
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
