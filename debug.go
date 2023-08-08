package main

import (
	"fmt"
)

func showEnv(e Expr, env *Env) Expr {
	if e.kind == SYMBOL {
		key := e.value.(string)
		value := lookup(env, key)
		printNode(value)

	} else {
		printEnv(env)
	}

	return Expr{NULL, nil, nil, nil}
}

func inspect(e Expr, env *Env) Expr {
	e = eval(&e, env)

	fmt.Printf("Expr: %v\n", e)
	fmt.Printf("Kind: %v\n", typeof(e.kind))
	fmt.Printf("Value: %v\n", e.value)
	fmt.Printf("Child: %v\n", e.child)
	fmt.Printf("Next: %v\n", e.next)

	printExpr(e.value.(*Expr))

	return Expr{NULL, nil, nil, nil}
}

func lookupValue(e Expr, env *Env) Expr {
	e = eval(&e, env)
	value := lookup(env, e.value.(string))
	return value
}
