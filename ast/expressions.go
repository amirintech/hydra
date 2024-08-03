package ast

import (
	"bytes"
	"strings"

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

// boolean literal
type Boolean struct {
	Token *token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

func (b *Boolean) String() string { return b.TokenLiteral() }

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

// if expression
type IfExpression struct {
	Token       *token.Token // if
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *IfExpression) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("if")
	buf.WriteString(ie.Condition.String())
	buf.WriteString(" ")
	buf.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		buf.WriteString("else ")
		buf.WriteString(ie.Alternative.String())
	}

	return buf.String()
}

// function literal
type FunctionLiteral struct {
	Token      *token.Token // fn
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

func (fl *FunctionLiteral) String() string {
	buf := bytes.Buffer{}
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	buf.WriteString(fl.TokenLiteral())
	buf.WriteString("(")
	buf.WriteString(strings.Join(params, ", "))
	buf.WriteString(") ")
	buf.WriteString(fl.Body.String())

	return buf.String()
}

// call expression
type CallExpression struct {
	Token     *token.Token // ( token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }

func (ce *CallExpression) String() string {
	buf := bytes.Buffer{}
	args := []string{}
	for _, arg := range ce.Arguments {
		args = append(args, arg.String())
	}

	buf.WriteString(ce.Function.String())
	buf.WriteString("(")
	buf.WriteString(strings.Join(args, ", "))
	buf.WriteString(")")

	return buf.String()
}
