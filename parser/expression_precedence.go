package parser

import "github.com/amirintech/hydra-compiler/token"

const (
	_ int = iota
	LOWEST
	EQUALS       // ==
	LESS_GREATER // < or >
	SUM          // +
	PRODUCT      // *
	PREFIX       // !x
	CALL         // foo()
)

var precedences = map[token.TokeType]int{
	token.EQUAL:        EQUALS,
	token.NOT_EQUAL:    EQUALS,
	token.LESS_THAN:    LESS_GREATER,
	token.GREATER_THAN: LESS_GREATER,
	token.PLUS:         SUM,
	token.MINUS:        SUM,
	token.SLASH:        PRODUCT,
	token.ASTERISK:     PRODUCT,
}

func (p *Parser) precedenceFromToken(t *token.Token) int {
	if p, ok := precedences[t.Type]; ok {
		return p
	}

	return LOWEST
}
