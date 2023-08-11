package env

import (
	"lisp/model"
)

func equals(operands []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoOperands(operands, env, evaluator)

	if a.Kind != b.Kind {
		return model.BoolExpr(false)
	}

	if a.Value == b.Value {
		return model.BoolExpr(true)
	} else {
		return model.BoolExpr(false)
	}
}

func notEquals(operands []*Expr, env *Env, evaluator Callback) *Expr {
	if equals(operands, env, evaluator).Value.(bool) {
		return model.BoolExpr(false)
	}

	return model.BoolExpr(true)
}

func lessThan(operands []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoOperands(operands, env, evaluator)

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("lessThan requires two numbers")
	}

	return model.BoolExpr(a.Value.(float64) < b.Value.(float64))
}

func greaterThan(operands []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoOperands(operands, env, evaluator)

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("greaterThan requires two numbers")
	}

	return model.BoolExpr(a.Value.(float64) > b.Value.(float64))
}

func lessThanEquals(operands []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoOperands(operands, env, evaluator)

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("lessThanEquals requires two numbers")
	}

	return model.BoolExpr(a.Value.(float64) <= b.Value.(float64))
}

func greaterThanEquals(operands []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoOperands(operands, env, evaluator)

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("greaterThanEquals requires two numbers")
	}

	return model.BoolExpr(a.Value.(float64) >= b.Value.(float64))
}
