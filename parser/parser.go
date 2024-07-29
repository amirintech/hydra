package parser

import (
	"fmt"

	"github.com/amirintech/hydra-compiler/ast"
	"github.com/amirintech/hydra-compiler/lexer"
	"github.com/amirintech/hydra-compiler/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	currToken *token.Token
	peekToken *token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer:  l,
		errors: []string{},
	}

	parser.nextToken() // init peekToken
	parser.nextToken() // init currToken

	return parser
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{Statements: []ast.Statement{}}

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}

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
		return nil
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
