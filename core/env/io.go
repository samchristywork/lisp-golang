package env

import (
	"fmt"
)

func __print(e *Expr, env *Env, evaluator Callback) {
	fmt.Print(evaluator(e, env).Value)
}

func _print(operands []*Expr, env *Env, evaluator Callback) *Expr {
	for _, operand := range operands {
		__print(operand, env, evaluator)
	}

	fmt.Println()

	return &Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}
}
