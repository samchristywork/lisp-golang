package main

import (
	"testing"
)

func TestBasicArithmetic(t *testing.T) {
	ret := evalNew(parse("(+ 1 2)"))

	if !ret.Equal(&Expr{NUMBER, 3.0, nil, nil}) {
		t.Error("Expected 3.0")
	}

	ret = evalNew(parse("(- -100 200)"))

	if !ret.Equal(&Expr{NUMBER, -300.0, nil, nil}) {
		t.Error("Expected -300.0")
	}

	ret = evalNew(parse("(* 2 2)"))

	if !ret.Equal(&Expr{NUMBER, 4.0, nil, nil}) {
		t.Error("Expected 4.0")
	}

	ret = evalNew(parse("(/ 8 2)"))

	if !ret.Equal(&Expr{NUMBER, 4.0, nil, nil}) {
		t.Error("Expected 4.0")
	}
}

func TestNestedArithmetic(t *testing.T) {
	ret := evalNew(parse("(+ 1 (+ 2 3))"))

	if !ret.Equal(&Expr{NUMBER, 6.0, nil, nil}) {
		t.Error("Expected 6.0")
	}
}
