package core

import (
	"fmt"
	"lisp/core/env"
	"lisp/model"
)

type Expr = model.Expr

func evalList(expr *Expr, env *core.Env) Expr {
	listChild := expr.Child

	if listChild == nil {
		return *expr
	}

	if listChild.Kind == core.LIST {
		return evalList(listChild, env)
	} else if listChild.Kind == core.SYMBOL && listChild.Value == "set" {
		key := listChild.Next.Value.(string)
		value := eval(listChild.Next.Next, env)

		core.AddEnv(env, key, value)

		return Expr{core.NULL, nil, nil, nil}
	} else if listChild.Kind == core.SYMBOL && listChild.Value == "define" {
		key := listChild.Next.Value.(string)

		core.AddEnv(env, key, Expr{core.PAMBDA, listChild, nil, nil})

		return Expr{core.NULL, nil, nil, nil}
	} else if listChild.Kind == core.SYMBOL && listChild.Value == "pambda" {
		body := listChild.Next.Next

		params := listChild.Next.Child
		next := expr.Next

		for {
			if params == nil {
				break
			}

			param1 := params.Value
			next1 := eval(next, env)
			core.AddEnv(env, param1.(string), next1)

			params = params.Next
			next = next.Next
		}

		result := eval(body, env)

		return eval(&result, env)
	} else if listChild.Kind == core.SYMBOL {

		// apply function
		key := listChild.Value.(string)
		value := env.Data[key]

		if value.Kind == core.FUNCTION {
			f := value.Value.(func(Expr, *core.Env, func(*Expr, *core.Env) Expr) Expr)

			if listChild.Next == nil {
				return f(core.Expr{core.NULL, nil, nil, nil}, env, eval)
			}

			return f(*listChild.Next, env, eval)
		} else if value.Kind == core.PAMBDA {
			substitution := core.Lookup(env, key).Value.(*Expr).Next.Next

			expr.Child = substitution
			expr.Child.Next = listChild.Next

			return eval(expr, env)
		} else if value.Kind == core.NULL {
			// Empty
		} else {
			fmt.Println(typeof(value.Kind))
			panic("evalList: unknown kind of symbol")
		}
	}

	return *expr
}

func evalAtom(expr *Expr, env *core.Env) Expr {
	if expr.Kind == core.SYMBOL {
		value := core.Lookup(env, expr.Value.(string))

		if value.Kind == core.UNKNOWN {
			fmt.Println(expr.Value.(string))
			panic("eval: unknown symbol")
		}

		return value
	} else if expr.Kind == core.NUMBER {
		return *expr
	} else if expr.Kind == core.STRING {
		return *expr
	} else if expr.Kind == core.PAMBDA {
		return *expr
	} else if expr.Kind == core.BOOL {
		return *expr
	} else if expr.Kind == core.NULL {
		return *expr
	}

	fmt.Println(typeof(expr.Kind))
	panic("eval: unknown kind")
}

func eval(expr *Expr, env *core.Env) Expr {
	result := Expr{core.UNKNOWN, nil, nil, nil}

	if expr.Kind == core.LIST {
		result = evalList(expr, env)
	} else {
		result = evalAtom(expr, env)
	}

	return result
}

func evalNew(expr *Expr) Expr {
	env := core.InitEnv()

	return eval(expr, env)
}
