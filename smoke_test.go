package main

import (
	"testing"
)

func TestSmoke(t *testing.T) {
	ret := evalNew(parse("(print \"Hello, World!\")"))

	if !ret.Equal(&Expr{NULL, nil, nil, nil}) {
		t.Error("Expected NULL")
	}
}

func TestBasicComparison(t *testing.T) {
	ret := evalNew(parse("(= 1 2)"))

	if ret.Kind != BOOL || ret.Value != false || ret.Child != nil || ret.Next != nil {
		t.Error("Expected false")
	}

	ret = evalNew(parse("(= 1 1)"))

	if !ret.Equal(&Expr{BOOL, true, nil, nil}) {
		t.Error("Expected true")
	}
}
