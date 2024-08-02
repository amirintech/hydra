package lexer

import (
	"github.com/amirintech/hydra-compiler/token"
)

type Lexer struct {
	input   string
	pos     int
	readPos int
	ch      byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar() // inits the lexer's state

	return l
}

func (l *Lexer) NextToken() *token.Token {
	var tok = &token.Token{}

	l.consumerWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok.Type = token.EQUAL
			tok.Literal = "=="
			l.readChar()
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			tok.Type = token.NOT_EQUAL
			tok.Literal = "!="
			l.readChar()
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case ',':
		tok = newToken(token.COMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LESS_THAN, l.ch)
	case '>':
		tok = newToken(token.GREATER_THAN, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok = &token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readSequence(isLetter)
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readSequence(isDigit)
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readSequence(predicate func(ch byte) bool) string {
	pos := l.pos
	for predicate(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}

	return l.input[l.readPos]
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}

	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) consumerWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
