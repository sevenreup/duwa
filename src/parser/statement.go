package parser

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/token"
)

func (p *Parser) parseStatement() ast.Statement {
	if p.curToken.Type == token.IDENT && p.peekTokenIs(token.ASSIGN) {
		return p.parseAssignmentStatement()
	}
	if token.LookupVariableType(p.curToken.Type) != "" {
		return p.parseVariableDeclarationStatement()
	}
	switch p.curToken.Type {
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
