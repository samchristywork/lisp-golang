package core

import (
	"fmt"
)

func __print(e Expr, env *Env) {
	fmt.Print(eval(&e, env).Value)

	if e.Next != nil {
		__print(*e.Next, env)
	}
}

func _print(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	__print(e, env)
	fmt.Println()

	return Expr{NULL, nil, nil, nil}
}
