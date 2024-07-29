package parser

import (
	"github.com/amirintech/hydra-compiler/ast"
	"github.com/amirintech/hydra-compiler/token"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.tokenIs(p.currToken, token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currToken}
	p.nextToken()

	for !p.tokenIs(p.currToken, token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
