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

func NewEnv(outer *Env) *Env {
	return &Env{outer, make(map[string]*Expr)}
}

func _lookup(env *Env, key string) *Expr {
	if env.Outer != nil {
		value := _lookup(env.Outer, key)
		if value != nil {
			return value
		}
	}

	value := env.Data[key]

	if value != nil {
		return value
	} else {
		return nil
	}
}

func Lookup(env *Env, key string) *Expr {
	if env == nil {
		fmt.Println("Environment not defined")
		os.Exit(1)
	}

	value := _lookup(env, key)

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

	// Data and Control Flow
	AddEnv(env, "assert", model.FunctionExpr(assert))
	AddEnv(env, "begin", model.FunctionExpr(begin))
	AddEnv(env, "define", model.FunctionExpr(define))
	AddEnv(env, "if", model.FunctionExpr(_if))
	AddEnv(env, "loop", model.FunctionExpr(loop))
	//AddEnv(env, "set", model.FunctionExpr(set))

	// Arithmetic
	AddEnv(env, "+", model.FunctionExpr(plus))
	AddEnv(env, "-", model.FunctionExpr(minus))
	AddEnv(env, "*", model.FunctionExpr(multiply))
	AddEnv(env, "/", model.FunctionExpr(divide))

	// Comparison
	AddEnv(env, "=", model.FunctionExpr(equals))
	AddEnv(env, "!=", model.FunctionExpr(notEquals))
	AddEnv(env, "<", model.FunctionExpr(lessThan))
	AddEnv(env, ">", model.FunctionExpr(greaterThan))
	AddEnv(env, "<=", model.FunctionExpr(lessThanEquals))
	AddEnv(env, ">=", model.FunctionExpr(greaterThanEquals))

	// Logic
	AddEnv(env, "and", model.FunctionExpr(and))
	AddEnv(env, "or", model.FunctionExpr(or))
	AddEnv(env, "not", model.FunctionExpr(not))
	AddEnv(env, "xor", model.FunctionExpr(xor))
	AddEnv(env, "nor", model.FunctionExpr(nor))
	AddEnv(env, "nand", model.FunctionExpr(nand))
	AddEnv(env, "xnor", model.FunctionExpr(xnor))

	// Constants
	AddEnv(env, "true", model.BoolExpr(true))
	AddEnv(env, "false", model.BoolExpr(false))
	AddEnv(env, "null", model.NullExpr())

	// I/O
	AddEnv(env, "print", model.FunctionExpr(_print))
	AddEnv(env, "println", model.FunctionExpr(_println))
	AddEnv(env, "system", model.FunctionExpr(system))

	// Debug
	AddEnv(env, "env", model.FunctionExpr(showEnv))
	//AddEnv(env, "inspect", model.FunctionExpr(inspect))
	//AddEnv(env, "lookup", model.FunctionExpr(lookupValue))

	return env
}
