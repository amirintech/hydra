package ast

import (
	"bytes"

	"github.com/amirintech/hydra-compiler/token"
)

// integer literal
type IntegerLiteral struct {
	Token *token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

func (il *IntegerLiteral) String() string { return il.TokenLiteral() }

// prefix expression
type PrefixExpression struct {
	Token      *token.Token
	Operator   string
	RightValue Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *PrefixExpression) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("(")
	buf.WriteString(pe.Operator)
	buf.WriteString(pe.RightValue.String())
	buf.WriteString(")")

	return buf.String()
}

// infix expression
type InfixExpression struct {
	Token      *token.Token
	LeftValue  Expression
	Operator   string
	RightValue Expression
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *InfixExpression) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("(")
	buf.WriteString(ie.LeftValue.String())
	buf.WriteString(" " + ie.Operator + " ")
	buf.WriteString(ie.RightValue.String())
	buf.WriteString(")")

	return buf.String()
}
