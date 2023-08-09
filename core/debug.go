package core

import (
	"fmt"
)

func showEnv(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	if e.Kind == SYMBOL {
		key := e.Value.(string)
		value := lookup(env, key)
		printNode(value)

	} else {
		printEnv(env)
	}

	return Expr{NULL, nil, nil, nil}
}

func inspect(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	e = eval(&e, env)

	fmt.Printf("Expr: %v\n", e)
	fmt.Printf("Kind: %v\n", typeof(e.Kind))
	fmt.Printf("Value: %v\n", e.Value)
	fmt.Printf("Child: %v\n", e.Child)
	fmt.Printf("Next: %v\n", e.Next)

	printExpr(e.Value.(*Expr))

	return Expr{NULL, nil, nil, nil}
}

func lookupValue(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	e = eval(&e, env)
	value := lookup(env, e.Value.(string))
	return value
}
