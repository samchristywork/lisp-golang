package env

import (
	"fmt"
)

func __print(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) {
	fmt.Print(evaluator(&e, env).Value)

	if e.Next != nil {
		__print(*e.Next, env, evaluator)
	}
}

func _print(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	__print(e, env, evaluator)
	fmt.Println()

	return Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}
}
