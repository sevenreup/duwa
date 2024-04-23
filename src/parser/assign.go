package parser

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/token"
)

func (p *Parser) parseAssignmentStatement() *ast.AssigmentStatement {
	statement := &ast.AssigmentStatement{}

	if p.curTokenIs(token.IDENT) {
		statement.Identifier = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	} else if p.curTokenIs(token.ASSIGN) {
		p.nextToken()
		indexExp := p.parseIndexExpression(statement.Identifier)
		statement.Identifier = indexExp

		p.nextToken()

		if !p.curTokenIs(token.ASSIGN) {
			return nil
		}

		p.nextToken()

		statement.Value = p.parseExpression(LOWEST)

		p.previousIndex = nil

		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}

		return statement
	}

	if !p.peekTokenIs(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	p.nextToken()

	statement.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}
