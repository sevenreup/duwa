package parser

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/token"
)

func (p *Parser) parseStatement() ast.Statement {
	if p.curToken.Type == token.RETURN {
		return p.parseReturnStatement()
	}

	if token.LookupVariableType(p.curToken.Type) != "" {
		return p.parseVariableDeclarationStatement()
	}

	statement := p.parseAssignmentStatement()

	if statement != nil {
		return statement
	}

	return p.parseExpressionStatement()
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	// if p.peekTokenIs(token.SEMICOLON) {
	// 	p.nextToken()
	// }
	return stmt
}
