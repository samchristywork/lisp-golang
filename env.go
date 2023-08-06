package main

import (
	"fmt"
)

type Env struct {
	outer *Env
	data  map[string]Expr
}

func addEnv(env *Env, key string, value Expr) {
	env.data[key] = value
}

func __print(e Expr, env *Env) {
	fmt.Print(eval(&e, env).value)
	if e.next != nil {
		__print(*e.next, env)
	}
}

func _print(e Expr, env *Env) Expr {
	__print(e, env)
	fmt.Println()

	return Expr{NULL, nil, nil, nil}
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

func define(e Expr, env *Env) Expr {
	key := e
	value := e.next

	if key.kind != SYMBOL {
		panic("define requires a symbol")
	}

	addEnv(env, key.value.(string), eval(value, env))

	return Expr{NULL, nil, nil, nil}
}

func _if(e Expr, env *Env) Expr {
	condition := eval(&e, env)

	if condition.value.(bool) {
		return eval(e.next, env)
	} else {
		return eval(e.next.next, env)
	}
}

func initEnv() *Env {
	env := &Env{nil, make(map[string]Expr)}
	// Control Flow
	addEnv(env, "begin", Expr{FUNCTION, begin, nil, nil})
	addEnv(env, "if", Expr{FUNCTION, _if, nil, nil})

	// Arithmetic
	addEnv(env, "+", Expr{FUNCTION, plus, nil, nil})
	addEnv(env, "-", Expr{FUNCTION, minus, nil, nil})
	addEnv(env, "*", Expr{FUNCTION, multiply, nil, nil})
	addEnv(env, "/", Expr{FUNCTION, divide, nil, nil})

	// Comparison
	addEnv(env, "=", Expr{FUNCTION, equals, nil, nil})
	addEnv(env, "!=", Expr{FUNCTION, notEquals, nil, nil})
	addEnv(env, "<", Expr{FUNCTION, lessThan, nil, nil})
	addEnv(env, ">", Expr{FUNCTION, greaterThan, nil, nil})
	addEnv(env, "<=", Expr{FUNCTION, lessThanEquals, nil, nil})
	addEnv(env, ">=", Expr{FUNCTION, greaterThanEquals, nil, nil})

	// Logic
	addEnv(env, "and", Expr{FUNCTION, and, nil, nil})
	addEnv(env, "or", Expr{FUNCTION, or, nil, nil})
	addEnv(env, "not", Expr{FUNCTION, not, nil, nil})
	addEnv(env, "xor", Expr{FUNCTION, xor, nil, nil})
	addEnv(env, "nor", Expr{FUNCTION, nor, nil, nil})
	addEnv(env, "nand", Expr{FUNCTION, nand, nil, nil})
	addEnv(env, "xnor", Expr{FUNCTION, xnor, nil, nil})

	// Constants
	addEnv(env, "false", Expr{BOOL, false, nil, nil})
	addEnv(env, "true", Expr{BOOL, true, nil, nil})
	addEnv(env, "null", Expr{NULL, nil, nil, nil})

	// Variables
	addEnv(env, "define", Expr{FUNCTION, define, nil, nil})

	// I/O
	addEnv(env, "print", Expr{FUNCTION, _print, nil, nil})
	return env
}
