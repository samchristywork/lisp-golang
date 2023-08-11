package env

import (
	"lisp/model"
)

func define(operands []*Expr, env *Env, evaluator Callback) *Expr {
	if len(operands) != 2 {
		panic("define requires 2 operands")
	}

	key := operands[0]
	value := operands[1]

	if key.Kind == model.SYMBOL {
		AddEnv(env, key.Value.(string), value)
	} else if key.Kind == model.LIST {
		l := &Expr{Kind: model.LAMBDA, Value: key, Next: nil, Child: value}
		AddEnv(env, key.Child.Value.(string), l)
	} else {
		panic("define requires a symbol")
	}

	return model.NullExpr()
}
