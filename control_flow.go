package main

import (
	"fmt"
	"os"
)

func _if(e Expr, env *Env) Expr {
	condition := eval(&e, env)

	if condition.Value.(bool) { // Consequent
		e.Next.Next = nil
		return eval(e.Next, env)

	} else { // Alternative
		return eval(e.Next.Next, env)
	}
}

func begin(e Expr, env *Env) Expr {
	ret := Expr{NULL, nil, nil, nil}

	for {
		ret = eval(&e, env)

		if e.Next == nil {
			break
		}

		e = *e.Next
	}

	return ret
}

func loop(e Expr, env *Env) Expr {
	ret := Expr{NULL, nil, nil, nil}

	head := e

	for {
		ret = eval(&e, env)

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

func assert(e Expr, env *Env) Expr {
	assertion := eval(&e, env)

	if assertion.Kind != BOOL {
		panic("assertion must be a boolean")
	}

	if assertion.Value.(bool) {
		return Expr{NULL, nil, nil, nil}

	} else {
		fmt.Printf("Assertion failed: %v\n", e)
		os.Exit(1)
		return Expr{NULL, nil, nil, nil}
	}
}
