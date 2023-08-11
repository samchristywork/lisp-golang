package env

import (
	"lisp/model"
)

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
