package lexer

import "github.com/amirintech/hydra-compiler/token"

func newToken(tokenType token.TokeType, literal byte) *token.Token {
	return &token.Token{
		Type:    tokenType,
		Literal: string(literal),
	}
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch == '_')
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
