package env

import (
	"lisp/model"
)

func expectTwoNumbers(operands []*Expr, env *Env, evaluator Callback) (*Expr, *Expr) {
	if len(operands) != 2 {
		panic("requires two operands")
	}

	a := evaluator(operands[0], env)
	b := evaluator(operands[1], env)

	if a.Kind != model.NUMBER || b.Kind != model.NUMBER {
		panic("requires two numbers")
	}

	return a, b
}

func expectTwoOperands(operands []*Expr, env *Env, evaluator Callback) (*Expr, *Expr) {
	if len(operands) != 2 {
		panic("requires two operands")
	}

	a := evaluator(operands[0], env)
	b := evaluator(operands[1], env)

	return a, b
}

func expectTwoBools(operands []*Expr, env *Env, evaluator Callback) (*Expr, *Expr) {
	if len(operands) != 2 {
		panic("requires two operands")
	}

	a := evaluator(operands[0], env)
	b := evaluator(operands[1], env)

	if a.Kind != model.BOOL || b.Kind != model.BOOL {
		panic("requires two booleans")
	}

	return a, b
}

func expectOneBool(operands []*Expr, env *Env, evaluator Callback) *Expr {
	if len(operands) != 1 {
		panic("requires one operand")
	}

	a := evaluator(operands[0], env)

	if a.Kind != model.BOOL {
		panic("requires one boolean")
	}

	return a
}
