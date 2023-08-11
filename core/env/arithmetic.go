package env

import (
	"lisp/model"
)

func expectTwoNumbers(operands []*Expr, env *Env, evaluator Callback) (*Expr, *Expr) {
	if len(operands) != 2 {
		panic("requires two operands")
	}

	a := evaluator(operands[0], env)
	b := evaluator(operands[1], env)

	if operands[0].Kind != NUMBER || operands[1].Kind != NUMBER {
		panic("requires two numbers")
	}

	return a, b
}

func plus(operands []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoNumbers(operands, env, evaluator)

	return model.NumberExpr(a.Value.(float64) + b.Value.(float64))
}

func minus(operands []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoNumbers(operands, env, evaluator)

	return model.NumberExpr(a.Value.(float64) - b.Value.(float64))
}

func multiply(operands []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoNumbers(operands, env, evaluator)

	return model.NumberExpr(a.Value.(float64) * b.Value.(float64))
}

func divide(operands []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoNumbers(operands, env, evaluator)

	return model.NumberExpr(a.Value.(float64) / b.Value.(float64))
}
