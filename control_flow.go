package main

import (
	"fmt"
	"os"
)

func _if(e Expr, env *Env) Expr {
	condition := eval(&e, env)

	if condition.value.(bool) { // Consequent
		e.next.next = nil
		return eval(e.next, env)

	} else { // Alternative
		return eval(e.next.next, env)
	}
}

func begin(e Expr, env *Env) Expr {
	ret := Expr{NULL, nil, nil, nil}

	for {
		ret = eval(&e, env)

		if e.next == nil {
			break
		}

		e = *e.next
	}

	return ret
}

func loop(e Expr, env *Env) Expr {
	ret := Expr{NULL, nil, nil, nil}

	head := e

	for {
		ret = eval(&e, env)

		if e.next == nil {
			if ret.kind == BOOL {
				if ret.value.(bool) {
				} else {
					break
				}
			} else {
				panic("loop requires a boolean expression")
			}

			e = head
			continue
		}

		e = *e.next
	}

	return ret
}

func assert(e Expr, env *Env) Expr {
	assertion := eval(&e, env)

	if assertion.kind != BOOL {
		panic("assertion must be a boolean")
	}

	if assertion.value.(bool) {
		return Expr{NULL, nil, nil, nil}

	} else {
		fmt.Printf("Assertion failed: %v\n", e)
		os.Exit(1)
		return Expr{NULL, nil, nil, nil}
	}
}
