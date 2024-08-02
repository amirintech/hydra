package parser

import (
	"fmt"

	"github.com/amirintech/hydra-compiler/ast"
	"github.com/amirintech/hydra-compiler/token"
)

func (p *Parser) peekError(t token.TokeType) {
	msg := fmt.Sprintf("expected next token to be of type '%s', but got type '%s' instead.", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) expectPeek(t token.TokeType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

func (p *Parser) tokenIs(tok *token.Token, t token.TokeType) bool {
	return tok.Type == t
}
