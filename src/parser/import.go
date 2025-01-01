package parser

import (
	"github.com/sevenreup/duwa/src/ast"
	"github.com/sevenreup/duwa/src/token"
)

func (p *Parser) importStatement() ast.Expression {
	statement := &ast.ImportExpression{Token: p.curToken}

	p.nextToken()

	if !p.expectPeek(token.STRING) {
		return nil
	}

	statement.Path = &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}

	return statement
}
