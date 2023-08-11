package env

import (
	"fmt"
	"os"
)

func _if(e *Expr, env *Env, evaluator Callback) *Expr {
	condition := evaluator(e, env)

	if condition.Value.(bool) { // Consequent
		e.Next.Next = nil
		return evaluator(e.Next, env)

	} else { // Alternative
		return evaluator(e.Next.Next, env)
	}
}

func begin(operands []*Expr, env *Env, evaluator Callback) *Expr {
	ret := &Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}

	for _, operand := range operands {
		ret = evaluator(operand, env)
	}

	return ret
}

func loop(e *Expr, env *Env, evaluator Callback) *Expr {
	ret := &Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}

	head := e

	for {
		ret = evaluator(e, env)

		if e.Next == nil {
			if ret.Kind == BOOL {
				if ret.Value.(bool) {
				} else {
					break
				}
			} else {
				panic("loop requires a boolean expression")
			}

			e = head
			continue
		}

		e = e.Next
	}

	return ret
}

func assert(operands []*Expr, env *Env, evaluator Callback) *Expr {
	if len(operands) != 1 {
		panic("assert requires one argument")
	}

	assertion := evaluator(operands[0], env)

	if assertion.Kind != BOOL {
		panic("assertion must be a boolean")
	}

	if assertion.Value.(bool) {
		return &Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}

	} else {
		fmt.Printf("Assertion failed.\n")
		os.Exit(1)
		return &Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}
	}
}
