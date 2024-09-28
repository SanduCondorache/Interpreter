package ast

import (
	"testing"

	"github.com/SanduCondorache/Interpreter/internals/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Id{
					Token: token.Token{Type: token.ID, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Id{
					Token: token.Token{Type: token.ID, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got = %q", program.String())
	}
}
