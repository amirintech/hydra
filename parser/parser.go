package parser

import (
	"github.com/amirintech/hydra-compiler/ast"
	"github.com/amirintech/hydra-compiler/lexer"
	"github.com/amirintech/hydra-compiler/token"
)

// types
type (
	prefixParseFunc func() ast.Expression
	infixParseFunc  func(ast.Expression) ast.Expression
)

type Parser struct {
	lexer            *lexer.Lexer
	currToken        *token.Token
	peekToken        *token.Token
	errors           []string
	prefixParseFuncs map[token.TokeType]prefixParseFunc
	infixParseFuncs  map[token.TokeType]infixParseFunc
}

// functions
func New(l *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer:            l,
		errors:           []string{},
		prefixParseFuncs: map[token.TokeType]prefixParseFunc{},
		infixParseFuncs:  map[token.TokeType]infixParseFunc{},
	}

	parser.nextToken() // init peekToken
	parser.nextToken() // init currToken

	// prefix functions
	prefixTokenTypes := []token.TokeType{
		token.IDENT, token.INT, token.BANG, token.MINUS,
	}
	prefixTokenFuncs := []prefixParseFunc{
		parser.parseIdentifier, parser.parseIntegerLiteral, parser.parsePrefixExpression, parser.parsePrefixExpression,
	}
	for i, tokType := range prefixTokenTypes {
		parser.registerPrefixFn(tokType, prefixTokenFuncs[i])
	}

	// infix functions
	infixTokenTypes := []token.TokeType{
		token.PLUS, token.MINUS, token.SLASH, token.ASTERISK,
		token.EQUAL, token.NOT_EQUAL, token.LESS_THAN, token.GREATER_THAN,
	}
	for _, tokType := range infixTokenTypes {
		parser.registerInfixFn(tokType, parser.parseInfixExpression)
	}

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
