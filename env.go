package main

import (
	"fmt"
	"os"
	"os/exec"
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

func set(e Expr, env *Env) Expr {
	key := eval(&e, env)
	value := eval(e.next, env)

	if key.kind != SYMBOL {
		panic("set requires a symbol")
	}

	addEnv(env, key.value.(string), eval(&value, env))

	return Expr{NULL, nil, nil, nil}
}

func define(e Expr, env *Env) Expr {
	key := eval(&e, env)
	value := eval(e.next, env)

	if key.kind != SYMBOL {
		panic("define requires a symbol")
	}

	addEnv(env, key.value.(string), eval(&value, env))

	return Expr{NULL, nil, nil, nil}
}

func _if(e Expr, env *Env) Expr {
	condition := eval(&e, env)

	if condition.value.(bool) { // Consequent
		e.next.next = nil
		return eval(e.next, env)

	} else { // Alternative
		return eval(e.next.next, env)
	}
}

func inspect(e Expr, env *Env) Expr {
	e = eval(&e, env)

	fmt.Printf("Expr: %v\n", e)
	fmt.Printf("Kind: %v\n", typeof(e.kind))
	fmt.Printf("Value: %v\n", e.value)
	fmt.Printf("Child: %v\n", e.child)
	fmt.Printf("Next: %v\n", e.next)

	printExpr(e.value.(*Expr))

	return Expr{NULL, nil, nil, nil}
}

func lookupValue(e Expr, env *Env) Expr {
	e = eval(&e, env)
	value := lookup(env, e.value.(string))
	return value
}

func lambda(e Expr, env *Env) Expr {
	fmt.Println("lambda is deprecated, use pambda instead.")

	return Expr{NULL, nil, nil, nil}
}

func pambda(e Expr, env *Env) Expr {
	return Expr{PAMBDA, e, nil, nil}
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

func showEnv(e Expr, env *Env) Expr {
	if e.kind == SYMBOL {
		key := e.value.(string)
		value := lookup(env, key)
		printNode(value)

	} else {
		printEnv(env)
	}

	return Expr{NULL, nil, nil, nil}
}

func system(e Expr, env *Env) Expr {
	e = eval(&e, env)

	if e.kind != STRING {
		panic("system requires a string")
	}

	args := []string{}

	head := e
	for {
		if head.kind != STRING {
			panic("system requires a string")
		}

		args = append(args, head.value.(string))

		if head.next == nil {
			break
		}

		head = *head.next
	}

	command := exec.Command(e.value.(string), args[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()

	return Expr{NULL, nil, nil, nil}
}

func initEnv() *Env {
	env := &Env{nil, make(map[string]Expr)}

	// Data and Control Flow
	addEnv(env, "assert", Expr{FUNCTION, assert, nil, nil})
	addEnv(env, "begin", Expr{FUNCTION, begin, nil, nil})
	addEnv(env, "define", Expr{FUNCTION, define, nil, nil})
	addEnv(env, "if", Expr{FUNCTION, _if, nil, nil})
	addEnv(env, "lambda", Expr{FUNCTION, lambda, nil, nil})
	addEnv(env, "loop", Expr{FUNCTION, loop, nil, nil})
	addEnv(env, "pambda", Expr{FUNCTION, pambda, nil, nil})
	addEnv(env, "set", Expr{FUNCTION, set, nil, nil})

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
	addEnv(env, "true", Expr{BOOL, true, nil, nil})
	addEnv(env, "false", Expr{BOOL, false, nil, nil})
	addEnv(env, "null", Expr{NULL, nil, nil, nil})

	// I/O
	addEnv(env, "print", Expr{FUNCTION, _print, nil, nil})
	addEnv(env, "system", Expr{FUNCTION, system, nil, nil})

	// Debug
	addEnv(env, "env", Expr{FUNCTION, showEnv, nil, nil})
	addEnv(env, "inspect", Expr{FUNCTION, inspect, nil, nil})
	addEnv(env, "lookup", Expr{FUNCTION, lookupValue, nil, nil})

	return env
}
