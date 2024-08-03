package ast

import (
	"bytes"

	"github.com/amirintech/hydra-compiler/token"
)

// identifier
type Identifier struct {
	Token *token.Token // IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }

// let statement
type LetStatement struct {
	Token *token.Token // LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	buf := bytes.Buffer{}

	buf.WriteString(ls.TokenLiteral() + " ")
	buf.WriteString(ls.Name.String())
	buf.WriteString(" = ")
	if ls.Value != nil {
		buf.WriteString(ls.Value.String())
	}
	buf.WriteString(";")

	return buf.String()
}

// return statement
type ReturnStatement struct {
	Token *token.Token
	Value Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	buf := bytes.Buffer{}

	buf.WriteString(rs.TokenLiteral() + " ")
	if rs.Value != nil {
		buf.WriteString(rs.Value.String())
	}
	buf.WriteString(";")

	return buf.String()
}

// expression statement
type ExpressionStatement struct {
	Token      *token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// block statement
type BlockStatement struct {
	Token      *token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

func (bs *BlockStatement) String() string {
	buf := bytes.Buffer{}
	for _, stmt := range bs.Statements {
		buf.WriteString(stmt.String())
	}

	return buf.String()
}
