package parser

import "github.com/sevenreup/duwa/src/ast"

func (p *Parser) nullLiteral() ast.Expression {
	return &ast.NullLiteral{Token: p.curToken}
}
