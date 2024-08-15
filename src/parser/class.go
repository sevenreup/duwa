package parser

import (
	"github.com/sevenreup/chewa/src/ast"
	"github.com/sevenreup/chewa/src/token"
)

func (p *Parser) classStatement() ast.Expression {
	class := &ast.ClassStatement{Token: p.curToken}

	p.nextToken()

	class.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// TODO: Implement inheritance

	if !p.expectPeek(token.OPENING_BRACE) {
		return nil
	}

	class.Body = p.parseBlockStatement()

	return class
}
