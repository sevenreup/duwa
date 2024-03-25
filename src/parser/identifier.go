package parser

import "github.com/sevenreup/chewa/src/ast"

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}
