package main

import (
	"fmt"
	core "lisp/core"
	environment "lisp/env"
	model "lisp/model"
	util "lisp/util"
	"testing"
)

func testExpression(t *testing.T, expression string, expected *model.Expr) {
	fmt.Println("Test: " + expression)

	ret := core.EvalNew(core.Parse(expression))

	if !ret.Equal(expected) {
		util.Red()
		fmt.Println("Test failed")
		environment.PrintExpr(ret)
		util.Reset()
		t.Error()
	} else {
		util.Green()
		fmt.Println("Test passed")
		util.Reset()
	}
}

func Test1(t *testing.T) {
	input := "(+ 1 2)"
	testExpression(t, input, model.NumberExpr(3.0))
}

func Test2(t *testing.T) {
	input := "(+ 1 (+ 1 2))"
	testExpression(t, input, model.NumberExpr(4.0))
}

func Test3(t *testing.T) {
	input := "(define x 1)"
	testExpression(t, input, model.NullExpr())
}

func Test4(t *testing.T) {
	input := "(assert true)"
	testExpression(t, input, model.NullExpr())
}

func Test5(t *testing.T) {
	input := "(begin (define x 1)(env x)(assert (= x 1)))"
	testExpression(t, input, model.NullExpr())
}

func Test6(t *testing.T) {
	input := "(begin (define x 1)(env x)(assert (= x 1)))"
	testExpression(t, input, model.NullExpr())
}

func Test7(t *testing.T) {
	input := "(begin (define x 1) x)"
	testExpression(t, input, model.NumberExpr(1.0))
}

func Test8(t *testing.T) {
	input := "(begin (print 1)(print 1 2 3 4 \"hi\"))"
	testExpression(t, input, model.NullExpr())
}

func Test9(t *testing.T) {
	input := "(begin (define (foo x) (+ 1 x)) (foo 2))"
	testExpression(t, input, model.NumberExpr(3.0))
}

func Test10(t *testing.T) {
	input := "(begin (define (print-hello-world) (print \"Hello, World!\")) (print-hello-world))"
	testExpression(t, input, model.NullExpr())
}

func Test11(t *testing.T) {
	input := "(begin (define (add-three-numbers x y z) (+ x (+ y z)))"
	input += "(add-three-numbers 1 2 3))"
	testExpression(t, input, model.NumberExpr(6.0))
}

func TestBasicArithmetic(t *testing.T) {
	input := "(+ 1 2)"
	testExpression(t, input, model.NumberExpr(3.0))

	input = "(- 1 2)"
	testExpression(t, input, model.NumberExpr(-1.0))

	input = "(* 1 2)"
	testExpression(t, input, model.NumberExpr(2.0))

	input = "(/ 1 2)"
	testExpression(t, input, model.NumberExpr(0.5))
}

func TestNestedArithmetic(t *testing.T) {
	input := "(+ 1 (+ 2 3))"
	testExpression(t, input, model.NumberExpr(6.0))
}

func TestSmoke(t *testing.T) {
	input := "(print \"Hello, World!\")"
	testExpression(t, input, model.NullExpr())
}

func TestComparison(t *testing.T) {
	input := "(= 1 1)"
	testExpression(t, input, model.BoolExpr(true))

	input = "(= 1 2)"
	testExpression(t, input, model.BoolExpr(false))

	input = "(= 1 (+ 1 1))"
	testExpression(t, input, model.BoolExpr(false))

	input = "(= 1 (+ 1 0))"
	testExpression(t, input, model.BoolExpr(true))

	input = "(!= 1 1)"
	testExpression(t, input, model.BoolExpr(false))

	input = "(!= 1 2)"
	testExpression(t, input, model.BoolExpr(true))

	input = "(< 1 2)"
	testExpression(t, input, model.BoolExpr(true))

	input = "(< 2 1)"
	testExpression(t, input, model.BoolExpr(false))

	input = "(< 1 1)"
	testExpression(t, input, model.BoolExpr(false))

	input = "(<= 1 2)"
	testExpression(t, input, model.BoolExpr(true))

	input = "(<= 2 1)"
	testExpression(t, input, model.BoolExpr(false))

	input = "(<= 1 1)"
	testExpression(t, input, model.BoolExpr(true))

	input = "(> 1 2)"
	testExpression(t, input, model.BoolExpr(false))

	input = "(> 2 1)"
	testExpression(t, input, model.BoolExpr(true))

	input = "(> 1 1)"
	testExpression(t, input, model.BoolExpr(false))

	input = "(>= 1 2)"
	testExpression(t, input, model.BoolExpr(false))

	input = "(>= 2 1)"
	testExpression(t, input, model.BoolExpr(true))

	input = "(>= 1 1)"
	testExpression(t, input, model.BoolExpr(true))
}
