package core

import (
	"fmt"
	"os"
)

func _if(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	condition := evaluator(&e, env)

	if condition.Value.(bool) { // Consequent
		e.Next.Next = nil
		return evaluator(e.Next, env)

	} else { // Alternative
		return evaluator(e.Next.Next, env)
	}
}

func begin(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	ret := Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}

	for {
		ret = evaluator(&e, env)

		if e.Next == nil {
			break
		}

		e = *e.Next
	}

	return ret
}

func loop(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	ret := Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}

	head := e

	for {
		ret = evaluator(&e, env)

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

		e = *e.Next
	}

	return ret
}

func assert(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	assertion := evaluator(&e, env)

	if assertion.Kind != BOOL {
		panic("assertion must be a boolean")
	}

	if assertion.Value.(bool) {
		return Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}

	} else {
		fmt.Printf("Assertion failed: %v\n", e)
		os.Exit(1)
		return Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}
	}
}
