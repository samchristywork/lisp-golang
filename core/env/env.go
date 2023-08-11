package env

import (
	"fmt"
	"lisp/model"
	"os"
	"os/exec"
)

type Expr = model.Expr
type Callback func(*Expr, *Env) *Expr

type Env struct {
	Outer *Env
	Data  map[string]*Expr
}

func AddEnv(env *Env, key string, value *Expr) {
	env.Data[key] = value
}

func Lookup(env *Env, key string) *Expr {
	if env == nil {
		fmt.Println("Environment not defined")
		os.Exit(1)
	}

	value := env.Data[key]

	if value == nil {
		fmt.Println("Unknown symbol: " + key)
		os.Exit(1)
	}

	return value
}

func system(operators []*Expr, env *Env, evaluator Callback) *Expr {
	args := []string{}

	if len(operators) < 1 {
		panic("system requires at least one argument")
	}

	for _, arg := range operators {
		e := evaluator(arg, env)

		if e.Kind != model.STRING {
			panic("system requires a string")
		}

		args = append(args, e.Value.(string))
	}

	command := exec.Command(args[0], args[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()

	return model.NullExpr()
}

func InitEnv() *Env {
	env := &Env{nil, make(map[string]*Expr)}

	//// Data and Control Flow
	AddEnv(env, "assert", &Expr{Kind: FUNCTION, Value: assert, Next: nil, Child: nil})
	AddEnv(env, "begin", &Expr{Kind: FUNCTION, Value: begin, Next: nil, Child: nil})
	AddEnv(env, "define", &Expr{Kind: FUNCTION, Value: define, Next: nil, Child: nil})
	AddEnv(env, "if", &Expr{Kind: FUNCTION, Value: _if, Next: nil, Child: nil})
	//AddEnv(env, "lambda", &Expr{Kind: FUNCTION, Value: lambda, Next: nil, Child: nil})
	AddEnv(env, "loop", &Expr{Kind: FUNCTION, Value: loop, Next: nil, Child: nil})
	//AddEnv(env, "pambda", &Expr{Kind: FUNCTION, Value: pambda, Next: nil, Child: nil})
	//AddEnv(env, "set", &Expr{Kind: FUNCTION, Value: set, Next: nil, Child: nil})

	// Arithmetic
	AddEnv(env, "+", &Expr{Kind: FUNCTION, Value: plus, Next: nil, Child: nil})
	AddEnv(env, "-", &Expr{Kind: FUNCTION, Value: minus, Next: nil, Child: nil})
	AddEnv(env, "*", &Expr{Kind: FUNCTION, Value: multiply, Next: nil, Child: nil})
	AddEnv(env, "/", &Expr{Kind: FUNCTION, Value: divide, Next: nil, Child: nil})

	//// Comparison
	AddEnv(env, "=", &Expr{Kind: FUNCTION, Value: equals, Next: nil, Child: nil})
	AddEnv(env, "!=", &Expr{Kind: FUNCTION, Value: notEquals, Next: nil, Child: nil})
	AddEnv(env, "<", &Expr{Kind: FUNCTION, Value: lessThan, Next: nil, Child: nil})
	AddEnv(env, ">", &Expr{Kind: FUNCTION, Value: greaterThan, Next: nil, Child: nil})
	AddEnv(env, "<=", &Expr{Kind: FUNCTION, Value: lessThanEquals, Next: nil, Child: nil})
	AddEnv(env, ">=", &Expr{Kind: FUNCTION, Value: greaterThanEquals, Next: nil, Child: nil})

	//// Logic
	//AddEnv(env, "and", &Expr{Kind: FUNCTION, Value: and, Next: nil, Child: nil})
	//AddEnv(env, "or", &Expr{Kind: FUNCTION, Value: or, Next: nil, Child: nil})
	//AddEnv(env, "not", &Expr{Kind: FUNCTION, Value: not, Next: nil, Child: nil})
	//AddEnv(env, "xor", &Expr{Kind: FUNCTION, Value: xor, Next: nil, Child: nil})
	//AddEnv(env, "nor", &Expr{Kind: FUNCTION, Value: nor, Next: nil, Child: nil})
	//AddEnv(env, "nand", &Expr{Kind: FUNCTION, Value: nand, Next: nil, Child: nil})
	//AddEnv(env, "xnor", &Expr{Kind: FUNCTION, Value: xnor, Next: nil, Child: nil})

	//// Constants
	AddEnv(env, "true", &Expr{Kind: BOOL, Value: true, Next: nil, Child: nil})
	AddEnv(env, "false", &Expr{Kind: BOOL, Value: false, Next: nil, Child: nil})
	//AddEnv(env, "null", &Expr{Kind: NULL, Value: nil, Next: nil, Child: nil})

	//// I/O
	AddEnv(env, "print", &Expr{Kind: FUNCTION, Value: _print, Next: nil, Child: nil})
	//AddEnv(env, "system", &Expr{Kind: FUNCTION, Value: system, Next: nil, Child: nil})

	//// Debug
	AddEnv(env, "env", &Expr{Kind: FUNCTION, Value: showEnv, Next: nil, Child: nil})
	//AddEnv(env, "inspect", &Expr{Kind: FUNCTION, Value: inspect, Next: nil, Child: nil})
	//AddEnv(env, "lookup", &Expr{Kind: FUNCTION, Value: lookupValue, Next: nil, Child: nil})

	return env
}
