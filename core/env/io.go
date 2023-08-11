package env

import (
	"fmt"
	"lisp/model"
)

func __print(e *Expr, env *Env, evaluator Callback) {
	fmt.Print(evaluator(e, env).Value)
}

func _print(operands []*Expr, env *Env, evaluator Callback) *Expr {
	for _, operand := range operands {
		__print(operand, env, evaluator)
	}

	fmt.Println()

	return model.NullExpr()
}
