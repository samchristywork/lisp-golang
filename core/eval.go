package core

import (
	"fmt"
)

func evalList(expr *Expr, env *Env) Expr {
	listChild := expr.Child

	if listChild == nil {
		return *expr
	}

	if listChild.Kind == LIST {
		return evalList(listChild, env)
	} else if listChild.Kind == SYMBOL && listChild.Value == "set" {
		key := listChild.Next.Value.(string)
		value := eval(listChild.Next.Next, env)

		addEnv(env, key, value)

		return Expr{NULL, nil, nil, nil}
	} else if listChild.Kind == SYMBOL && listChild.Value == "define" {
		key := listChild.Next.Value.(string)

		addEnv(env, key, Expr{PAMBDA, listChild, nil, nil})

		return Expr{NULL, nil, nil, nil}
	} else if listChild.Kind == SYMBOL && listChild.Value == "pambda" {
		body := listChild.Next.Next

		params := listChild.Next.Child
		next := expr.Next

		for {
			if params == nil {
				break
			}

			param1 := params.Value
			next1 := eval(next, env)
			addEnv(env, param1.(string), next1)

			params = params.Next
			next = next.Next
		}

		result := eval(body, env)

		return eval(&result, env)
	} else if listChild.Kind == SYMBOL {

		// apply function
		key := listChild.Value.(string)
		value := env.data[key]

		if value.Kind == FUNCTION {
			f := value.Value.(func(Expr, *Env, func(*Expr, *Env) Expr) Expr)

			if listChild.Next == nil {
				return f(Expr{NULL, nil, nil, nil}, env, eval)
			}

			return f(*listChild.Next, env, eval)
		} else if value.Kind == PAMBDA {
			substitution := lookup(env, key).Value.(*Expr).Next.Next

			expr.Child = substitution
			expr.Child.Next = listChild.Next

			return eval(expr, env)
		} else if value.Kind == NULL {
			// Empty
		} else {
			fmt.Println(typeof(value.Kind))
			panic("evalList: unknown kind of symbol")
		}
	}

	return *expr
}

func evalAtom(expr *Expr, env *Env) Expr {
	if expr.Kind == SYMBOL {
		value := lookup(env, expr.Value.(string))

		if value.Kind == UNKNOWN {
			fmt.Println(expr.Value.(string))
			panic("eval: unknown symbol")
		}

		return value
	} else if expr.Kind == NUMBER {
		return *expr
	} else if expr.Kind == STRING {
		return *expr
	} else if expr.Kind == PAMBDA {
		return *expr
	} else if expr.Kind == BOOL {
		return *expr
	} else if expr.Kind == NULL {
		return *expr
	}

	fmt.Println(typeof(expr.Kind))
	panic("eval: unknown kind")
}

func eval(expr *Expr, env *Env) Expr {
	result := Expr{UNKNOWN, nil, nil, nil}

	if expr.Kind == LIST {
		result = evalList(expr, env)
	} else {
		result = evalAtom(expr, env)
	}

	return result
}

func evalNew(expr *Expr) Expr {
	env := initEnv()

	return eval(expr, env)
}
