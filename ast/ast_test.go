package ast

import (
	"testing"

	"github.com/amirintech/hydra-compiler/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: &token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: &token.Token{Type: token.IDENT, Literal: "foo"},
					Value: "foo",
				},
				Value: &Identifier{
					Token: &token.Token{Type: token.IDENT, Literal: "bar"},
					Value: "bar",
				},
			},
		},
	}

	if program.String() != "let foo = bar;" {
		t.Errorf("program.String() returns a wrong value. got='%s'", program.String())
	}
}
