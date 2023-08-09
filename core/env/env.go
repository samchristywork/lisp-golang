package core

import (
	"fmt"
	"lisp/model"
	"os"
	"os/exec"
)

// TODO: Fix this
const (
	UNKNOWN = iota
	BOOL
	FUNCTION
	LAMBDA
	LIST
	NULL
	NUMBER
	PAMBDA
	STRING
	SYMBOL
)

type Expr = model.Expr

type Env struct {
	Outer *Env
	Data  map[string]Expr
}

func AddEnv(env *Env, key string, value Expr) {
	env.Data[key] = value
}

func Lookup(env *Env, key string) Expr {
	if env == nil {
		return Expr{Kind: UNKNOWN, Value: nil, Next: nil, Child: nil}
	}

	value := env.Data[key]

	if value.Kind == UNKNOWN {
		return Lookup(env.Outer, key)
	}

	return value
}

func lambda(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	fmt.Println("lambda is deprecated, use pambda instead.")

	return Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}
}

func pambda(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	return Expr{Kind: PAMBDA, Value: e, Next: nil, Child: nil}
}

func system(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	e = evaluator(&e, env)

	if e.Kind != STRING {
		panic("system requires a string")
	}

	args := []string{}

	head := e
	for {
		if head.Kind != STRING {
			panic("system requires a string")
		}

		args = append(args, head.Value.(string))

		if head.Next == nil {
			break
		}

		head = *head.Next
	}

	command := exec.Command(e.Value.(string), args[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()

	return Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}
}

func InitEnv() *Env {
	env := &Env{nil, make(map[string]Expr)}

	// Data and Control Flow
	AddEnv(env, "assert", Expr{Kind: FUNCTION, Value: assert, Next: nil, Child: nil})
	AddEnv(env, "begin", Expr{Kind: FUNCTION, Value: begin, Next: nil, Child: nil})
	AddEnv(env, "define", Expr{Kind: FUNCTION, Value: define, Next: nil, Child: nil})
	AddEnv(env, "if", Expr{Kind: FUNCTION, Value: _if, Next: nil, Child: nil})
	AddEnv(env, "lambda", Expr{Kind: FUNCTION, Value: lambda, Next: nil, Child: nil})
	AddEnv(env, "loop", Expr{Kind: FUNCTION, Value: loop, Next: nil, Child: nil})
	AddEnv(env, "pambda", Expr{Kind: FUNCTION, Value: pambda, Next: nil, Child: nil})
	AddEnv(env, "set", Expr{Kind: FUNCTION, Value: set, Next: nil, Child: nil})

	// Arithmetic
	AddEnv(env, "+", Expr{Kind: FUNCTION, Value: plus, Next: nil, Child: nil})
	AddEnv(env, "-", Expr{Kind: FUNCTION, Value: minus, Next: nil, Child: nil})
	AddEnv(env, "*", Expr{Kind: FUNCTION, Value: multiply, Next: nil, Child: nil})
	AddEnv(env, "/", Expr{Kind: FUNCTION, Value: divide, Next: nil, Child: nil})

	// Comparison
	AddEnv(env, "=", Expr{Kind: FUNCTION, Value: equals, Next: nil, Child: nil})
	AddEnv(env, "!=", Expr{Kind: FUNCTION, Value: notEquals, Next: nil, Child: nil})
	AddEnv(env, "<", Expr{Kind: FUNCTION, Value: lessThan, Next: nil, Child: nil})
	AddEnv(env, ">", Expr{Kind: FUNCTION, Value: greaterThan, Next: nil, Child: nil})
	AddEnv(env, "<=", Expr{Kind: FUNCTION, Value: lessThanEquals, Next: nil, Child: nil})
	AddEnv(env, ">=", Expr{Kind: FUNCTION, Value: greaterThanEquals, Next: nil, Child: nil})

	// Logic
	AddEnv(env, "and", Expr{Kind: FUNCTION, Value: and, Next: nil, Child: nil})
	AddEnv(env, "or", Expr{Kind: FUNCTION, Value: or, Next: nil, Child: nil})
	AddEnv(env, "not", Expr{Kind: FUNCTION, Value: not, Next: nil, Child: nil})
	AddEnv(env, "xor", Expr{Kind: FUNCTION, Value: xor, Next: nil, Child: nil})
	AddEnv(env, "nor", Expr{Kind: FUNCTION, Value: nor, Next: nil, Child: nil})
	AddEnv(env, "nand", Expr{Kind: FUNCTION, Value: nand, Next: nil, Child: nil})
	AddEnv(env, "xnor", Expr{Kind: FUNCTION, Value: xnor, Next: nil, Child: nil})

	// Constants
	AddEnv(env, "true", Expr{Kind: BOOL, Value: true, Next: nil, Child: nil})
	AddEnv(env, "false", Expr{Kind: BOOL, Value: false, Next: nil, Child: nil})
	AddEnv(env, "null", Expr{Kind: NULL, Value: nil, Next: nil, Child: nil})

	// I/O
	AddEnv(env, "print", Expr{Kind: FUNCTION, Value: _print, Next: nil, Child: nil})
	AddEnv(env, "system", Expr{Kind: FUNCTION, Value: system, Next: nil, Child: nil})

	// Debug
	AddEnv(env, "env", Expr{Kind: FUNCTION, Value: showEnv, Next: nil, Child: nil})
	AddEnv(env, "inspect", Expr{Kind: FUNCTION, Value: inspect, Next: nil, Child: nil})
	AddEnv(env, "lookup", Expr{Kind: FUNCTION, Value: lookupValue, Next: nil, Child: nil})

	return env
}
