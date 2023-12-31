package env

import (
	"fmt"
	"lisp/model"
	"os"
)

func _if(operands []*Expr, env *Env, evaluator Callback) *Expr {
	if len(operands) != 3 {
		panic("if requires three arguments")
	}

	condition := evaluator(operands[0], env)

	env = NewEnv(env)

	if condition.Value.(bool) { // Consequent
		return evaluator(operands[1], env)

	} else { // Alternative
		return evaluator(operands[2], env)
	}
}

func begin(operands []*Expr, env *Env, evaluator Callback) *Expr {
	ret := model.NullExpr()

	env = NewEnv(env)

	for _, operand := range operands {
		ret = evaluator(operand, env)
	}

	return ret
}

func loop(operands []*Expr, env *Env, evaluator Callback) *Expr {
	ret := model.NullExpr()

	env = NewEnv(env)

	for {
		for _, operand := range operands {
			ret = evaluator(operand, env)
		}

		if ret.Kind == model.BOOL {
			if !ret.Value.(bool) {
				break
			}
		} else {
			panic("loop requires a boolean expression")
		}
	}

	return ret
}

func assert(operands []*Expr, env *Env, evaluator Callback) *Expr {
	if len(operands) != 1 {
		panic("assert requires one argument")
	}

	assertion := evaluator(operands[0], env)

	if assertion.Kind != model.BOOL {
		panic("assertion must be a boolean")
	}

	if assertion.Value.(bool) {
		return model.NullExpr()
	} else {
		fmt.Printf("Assertion failed.\n")
		os.Exit(1)
		return model.NullExpr()
	}
}
