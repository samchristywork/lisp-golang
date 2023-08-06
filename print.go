package main

import (
	"fmt"
)

// TODO: printExpr is not used, but it is useful for debugging.
func _printExpr(expr *Expr, depth int) {
	if expr == nil {
		return
	}

	if expr.kind == LIST {
		fmt.Println()
	}

	if expr.kind == LIST {
		for i := 0; i < depth; i++ {
			fmt.Print("  ")
		}
		fmt.Print("( ")
		_printExpr(expr.child, depth+1)
		fmt.Print(" )")

	} else {
		fmt.Print(expr.value, " ", typeof(expr.kind))

		if (expr.next != nil) && (expr.next.kind != LIST) {
			fmt.Println()
		}
		if expr.next != nil {
			for i := 0; i < depth; i++ {
				fmt.Print("  ")
			}
		}
	}

	if expr.next != nil {
		_printExpr(expr.next, depth)
	}
}

func printExpr(expr *Expr) {
	_printExpr(expr, 0)
	fmt.Println()
}

func printEnv(env *Env) {
	fmt.Println("env:")
	for env != nil {
		for key, value := range env.data {
			fmt.Printf("%s\t%v\n", key, value.value)
		}
		env = env.outer
	}
}
