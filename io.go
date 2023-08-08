package main

import (
	"fmt"
)

func __print(e Expr, env *Env) {
	fmt.Print(eval(&e, env).value)

	if e.next != nil {
		__print(*e.next, env)
	}
}

func _print(e Expr, env *Env) Expr {
	__print(e, env)
	fmt.Println()

	return Expr{NULL, nil, nil, nil}
}
