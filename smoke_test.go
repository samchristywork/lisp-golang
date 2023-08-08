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

func TestBasicArithmetic(t *testing.T) {
	ret := evalNew(parse("(+ 1 2)"))

	if !ret.Equal(&Expr{NUMBER, 3.0, nil, nil}) {
		t.Error("Expected 3.0")
	}

	ret = evalNew(parse("(- -100 200)"))

	if !ret.Equal(&Expr{NUMBER, -300.0, nil, nil}) {
		t.Error("Expected -300.0")
	}
}

func TestNestedArithmetic(t *testing.T) {
	ret := evalNew(parse("(+ 1 (+ 2 3))"))

	if !ret.Equal(&Expr{NUMBER, 6.0, nil, nil}) {
		t.Error("Expected 6.0")
	}
}

func TestBasicComparison(t *testing.T) {
	ret := evalNew(parse("(= 1 2)"))

	if ret.kind != BOOL || ret.value != false || ret.child != nil || ret.next != nil {
		t.Error("Expected false")
	}

	ret = evalNew(parse("(= 1 1)"))

	if !ret.Equal(&Expr{BOOL, true, nil, nil}) {
		t.Error("Expected true")
	}
}
