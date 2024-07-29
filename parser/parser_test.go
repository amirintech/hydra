package parser

import (
	"testing"

	"github.com/amirintech/hydra-compiler/ast"
	"github.com/amirintech/hydra-compiler/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
   	let y = 10;
   	let foo = 838383;
	`
	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if program == nil {
		t.Fatal("ParseProgram() returned nil.\n")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statments contains %d statements instead of 3.", len(program.Statements))
	}

	tests := []struct{ expectedIdentifier string }{
		{"x"}, {"y"}, {"foo"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("Error: incorrect token literal.\texpected='let'\tgot='%s'", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("Error: incorrect statment.\texpected='*ast.LetStatment'\tgot='%T'", stmt)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("Error: incorrect name value.\texpected='%s'\tgot='%s'", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("Error: incorrect value token literal.\texpected='%s'\tgot='%s'", name, letStmt.Value.TokenLiteral())
		return false
	}

	return true
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 0;
	return 293317;
	`
	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statments contains %d statements instead of 3.", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Error: incorrect statment.\texpected='*ast.ReturnStatement'\tgot='%T'", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("Error: incorrect value token literal.\texpected='%s'\tgot='%s'", "return", returnStmt.TokenLiteral())
		}
	}

}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors.", len(errors))
	for i, msg := range errors {
		t.Errorf("parser error [%d]: %s", i+1, msg)
	}
	t.FailNow()
}
