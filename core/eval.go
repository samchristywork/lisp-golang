package core

import (
	"fmt"
	environment "lisp/core/env"
	"lisp/model"
)

type Expr = model.Expr

func evalList(expr *Expr, env *environment.Env) Expr {
	listChild := expr.Child

	if listChild == nil {
		return *expr
	}

	if listChild.Kind == environment.LIST {
		return evalList(listChild, env)
	} else if listChild.Kind == environment.SYMBOL && listChild.Value == "set" {
		key := listChild.Next.Value.(string)
		value := eval(listChild.Next.Next, env)

		environment.AddEnv(env, key, value)

		return Expr{Kind: environment.NULL, Value: nil, Next: nil, Child: nil}
	} else if listChild.Kind == environment.SYMBOL && listChild.Value == "define" {
		key := listChild.Next.Value.(string)

		environment.AddEnv(env, key, Expr{Kind: environment.PAMBDA, Value: listChild, Next: nil, Child: nil})

		return Expr{Kind: environment.NULL, Value: nil, Next: nil, Child: nil}
	} else if listChild.Kind == environment.SYMBOL && listChild.Value == "pambda" {
		body := listChild.Next.Next

		params := listChild.Next.Child
		next := expr.Next

		for {
			if params == nil {
				break
			}

			param1 := params.Value
			next1 := eval(next, env)
			environment.AddEnv(env, param1.(string), next1)

			params = params.Next
			next = next.Next
		}

		result := eval(body, env)

		return eval(&result, env)
	} else if listChild.Kind == environment.SYMBOL {

		// apply function
		key := listChild.Value.(string)
		value := env.Data[key]

		if value.Kind == environment.FUNCTION {
			f := value.Value.(func(Expr, *environment.Env, func(*Expr, *environment.Env) Expr) Expr)

			if listChild.Next == nil {
				return f(environment.Expr{Kind: environment.NULL, Value: nil, Next: nil, Child: nil}, env, eval)
			}

			return f(*listChild.Next, env, eval)
		} else if value.Kind == environment.PAMBDA {
			substitution := environment.Lookup(env, key).Value.(*Expr).Next.Next

			expr.Child = substitution
			expr.Child.Next = listChild.Next

			return eval(expr, env)
		} else if value.Kind == environment.NULL {
			// Empty
		} else {
			fmt.Println(typeof(value.Kind))
			panic("evalList: unknown kind of symbol")
		}
	}

	return *expr
}

func evalAtom(expr *Expr, env *environment.Env) Expr {
	if expr.Kind == environment.SYMBOL {
		value := environment.Lookup(env, expr.Value.(string))

		if value.Kind == environment.UNKNOWN {
			fmt.Println(expr.Value.(string))
			panic("eval: unknown symbol")
		}

		return value
	} else if expr.Kind == environment.NUMBER {
		return *expr
	} else if expr.Kind == environment.STRING {
		return *expr
	} else if expr.Kind == environment.PAMBDA {
		return *expr
	} else if expr.Kind == environment.BOOL {
		return *expr
	} else if expr.Kind == environment.NULL {
		return *expr
	}

	fmt.Println(typeof(expr.Kind))
	panic("eval: unknown kind")
}

func eval(expr *Expr, env *environment.Env) Expr {
	result := Expr{Kind: environment.UNKNOWN, Value: nil, Next: nil, Child: nil}

	if expr.Kind == environment.LIST {
		result = evalList(expr, env)
	} else {
		result = evalAtom(expr, env)
	}

	return result
}

func evalNew(expr *Expr) Expr {
	env := environment.InitEnv()

	return eval(expr, env)
}
