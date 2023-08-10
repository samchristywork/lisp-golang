package core

import (
	environment "lisp/core/env"
	"testing"
)

func TestBasicArithmetic(t *testing.T) {
	ret := evalNew(parse("(+ 1 2)"))

	if !ret.Equal(&Expr{Kind: environment.NUMBER, Value: 3.0, Next: nil, Child: nil}) {
		t.Error("Expected 3.0")
	}

	ret = evalNew(parse("(- -100 200)"))

	if !ret.Equal(&Expr{Kind: environment.NUMBER, Value: -300.0, Next: nil, Child: nil}) {
		t.Error("Expected -300.0")
	}

	ret = evalNew(parse("(* 2 2)"))

	if !ret.Equal(&Expr{Kind: environment.NUMBER, Value: 4.0, Next: nil, Child: nil}) {
		t.Error("Expected 4.0")
	}

	ret = evalNew(parse("(/ 8 2)"))

	if !ret.Equal(&Expr{Kind: environment.NUMBER, Value: 4.0, Next: nil, Child: nil}) {
		t.Error("Expected 4.0")
	}
}

func TestNestedArithmetic(t *testing.T) {
	ret := evalNew(parse("(+ 1 (+ 2 3))"))

	if !ret.Equal(&Expr{Kind: environment.NUMBER, Value: 6.0, Next: nil, Child: nil}) {
		t.Error("Expected 6.0")
	}
}
