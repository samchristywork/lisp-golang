package core

import (
	environment "lisp/core/env"
	"testing"
)

func TestSmoke(t *testing.T) {
	ret := evalNew(parse("(print \"Hello, World!\")"))

	if !ret.Equal(&environment.Expr{Kind: environment.NULL, Value: nil, Next: nil, Child: nil}) {
		t.Error("Expected NULL")
	}
}

func TestBasicComparison(t *testing.T) {
	ret := evalNew(parse("(= 1 2)"))

	if ret.Kind != environment.BOOL || ret.Value != false || ret.Child != nil || ret.Next != nil {
		t.Error("Expected false")
	}

	ret = evalNew(parse("(= 1 1)"))

	if !ret.Equal(&Expr{Kind: environment.BOOL, Value: true, Next: nil, Child: nil}) {
		t.Error("Expected true")
	}
}
