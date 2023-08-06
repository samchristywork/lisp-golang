package main

import (
	"fmt"
)

func eval(expr *Expr, env *Env) Expr {
	if expr.kind == LIST {
		if expr.child.kind == SYMBOL {
			key := expr.child.value.(string)
			value := env.data[key]

			if value.kind == FUNCTION {
				f := value.value.(func(Expr, *Env) Expr)

				return f(*expr.child.next, env)
			} else {
				fmt.Printf("error: %s %s not a function\n", expr.child.value, typeof(expr.child.kind))
				return value
			}
		}
	}

	if expr.kind == SYMBOL {
		value := env.data[expr.value.(string)]

		return value
	}

	return *expr
}
