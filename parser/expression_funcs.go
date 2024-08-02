package parser

import (
	"fmt"
	"strconv"

	"github.com/amirintech/hydra-compiler/ast"
	"github.com/amirintech/hydra-compiler/token"
)

func (p *Parser) registerPrefixFn(tokType token.TokeType, fn prefixParseFunc) {
	p.prefixParseFuncs[tokType] = fn
}

func (p *Parser) registerInfixFn(tokType token.TokeType, fn infixParseFunc) {
	p.infixParseFuncs[tokType] = fn
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFuncs[p.currToken.Type]
	if prefix == nil {
		p.noPrefixParseFuncError(p.currToken.Type)
		return nil
	}

	leftExp := prefix()
	for !p.tokenIs(p.peekToken, token.SEMICOLON) && precedence < p.precedenceFromToken(p.peekToken) {
		infix := p.infixParseFuncs[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) noPrefixParseFuncError(t token.TokeType) {
	msg := fmt.Sprintf("no prefix parse function for %s found.", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	literal := &ast.IntegerLiteral{Token: p.currToken}

	value, err := strconv.ParseInt(p.currToken.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("error parsing %q as integer.",
			p.currToken.Literal))
		return nil
	}

	literal.Value = value

	return literal
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
	}

	p.nextToken()
	exp.RightValue = p.parseExpression(PREFIX)

	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:     p.currToken,
		LeftValue: left,
		Operator:  p.currToken.Literal,
	}

	precedence := p.precedenceFromToken(p.currToken)
	p.nextToken()
	exp.RightValue = p.parseExpression(precedence)

	return exp
}
