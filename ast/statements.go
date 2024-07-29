package ast

import (
	"github.com/amirintech/hydra-compiler/token"
)

// identifier
type Identifier struct {
	Token *token.Token // IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// let statement
type LetStatement struct {
	Token *token.Token // LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// return statement
type ReturnStatement struct {
	Token *token.Token
	Value Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
