package env

import (
	"lisp/model"
)

func and(operators []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoBools(operators, env, evaluator)

	return model.BoolExpr(a.Value.(bool) && b.Value.(bool))
}

func or(operators []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoBools(operators, env, evaluator)

	return model.BoolExpr(a.Value.(bool) || b.Value.(bool))
}

func not(operators []*Expr, env *Env, evaluator Callback) *Expr {
	a := expectOneBool(operators, env, evaluator)

	return model.BoolExpr(!a.Value.(bool))
}

func xor(operators []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoBools(operators, env, evaluator)

	left := a.Value.(bool) && !b.Value.(bool)
	right := !a.Value.(bool) && b.Value.(bool)

	return model.BoolExpr(left || right)
}

func nor(operators []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoBools(operators, env, evaluator)

	return model.BoolExpr(!(a.Value.(bool) || b.Value.(bool)))
}

func nand(operators []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoBools(operators, env, evaluator)

	return model.BoolExpr(!(a.Value.(bool) && b.Value.(bool)))
}

func xnor(operators []*Expr, env *Env, evaluator Callback) *Expr {
	a, b := expectTwoBools(operators, env, evaluator)

	left := a.Value.(bool) && !b.Value.(bool)
	right := !a.Value.(bool) && b.Value.(bool)

	return model.BoolExpr(!(left || right))
}
