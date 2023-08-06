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

func equals(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)

	if a.kind != b.kind {
		return Expr{BOOL, false, nil, nil}
	}

	if a.value == b.value {
		return Expr{BOOL, true, nil, nil}
	} else {
		return Expr{BOOL, false, nil, nil}
	}
}

func notEquals(e Expr, env *Env) Expr {
	if equals(e, env).value.(bool) {
		return Expr{BOOL, false, nil, nil}
	}
	return Expr{BOOL, true, nil, nil}
}

func plus(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)

	if a.kind != NUMBER || b.kind != NUMBER {
		panic("plus requires two numbers")
	}

	return Expr{NUMBER, a.value.(float64) + b.value.(float64), nil, nil}
}

func minus(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)

	if a.kind != NUMBER || b.kind != NUMBER {
		panic("minus requires two numbers")
	}

	return Expr{NUMBER, a.value.(float64) - b.value.(float64), nil, nil}
}

func multiply(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)

	if a.kind != NUMBER || b.kind != NUMBER {
		panic("multiply requires two numbers")
	}

	return Expr{NUMBER, a.value.(float64) * b.value.(float64), nil, nil}
}

func divide(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)

	if a.kind != NUMBER || b.kind != NUMBER {
		panic("divide requires two numbers")
	}

	return Expr{NUMBER, a.value.(float64) / b.value.(float64), nil, nil}
}

func _print(e Expr, env *Env) Expr {
	fmt.Println(eval(&e, env).value)
	if e.next != nil {
		_print(*e.next, env)
	}
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

func lessThan(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)

	if a.kind != NUMBER || b.kind != NUMBER {
		panic("lessThan requires two numbers")
	}

	return Expr{BOOL, a.value.(float64) < b.value.(float64), nil, nil}
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
	addEnv(env, "!=", Expr{FUNCTION, notEquals, nil, nil})
	addEnv(env, "<", Expr{FUNCTION, lessThan, nil, nil})
	addEnv(env, "=", Expr{FUNCTION, equals, nil, nil})

	// Constants
	addEnv(env, "false", Expr{BOOL, false, nil, nil})
	addEnv(env, "true", Expr{BOOL, true, nil, nil})

	// Variables
	addEnv(env, "define", Expr{FUNCTION, define, nil, nil})

	// I/O
	addEnv(env, "print", Expr{FUNCTION, _print, nil, nil})
	return env
}