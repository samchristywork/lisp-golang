package main

import (
	"fmt"
)

func evalList(expr *Expr, env *Env) Expr {
	listChild := expr.child

	if listChild == nil {
		return *expr
	}

	if listChild.kind == LIST {
		return evalList(listChild, env)
	} else if listChild.kind == SYMBOL && listChild.value == "set" {
		key := listChild.next.value.(string)
		value := eval(listChild.next.next, env)

		addEnv(env, key, value)

		return Expr{NULL, nil, nil, nil}
	} else if listChild.kind == SYMBOL && listChild.value == "define" {
		key := listChild.next.value.(string)

		addEnv(env, key, Expr{PAMBDA, listChild, nil, nil})

		return Expr{NULL, nil, nil, nil}
	} else if listChild.kind == SYMBOL && listChild.value == "pambda" {
		body := listChild.next.next

		params := listChild.next.child
		next := expr.next

		for {
			if params == nil {
				break
			}

			param1 := params.value
			next1 := eval(next, env)
			addEnv(env, param1.(string), next1)

			params = params.next
			next = next.next
		}

		result := eval(body, env)

		return eval(&result, env)
	} else if listChild.kind == SYMBOL {

		// apply function
		key := listChild.value.(string)
		value := env.data[key]

		if value.kind == FUNCTION {
			f := value.value.(func(Expr, *Env) Expr)

			if listChild.next == nil {
				return f(Expr{NULL, nil, nil, nil}, env)
			}

			return f(*listChild.next, env)
		} else if value.kind == PAMBDA {
			substitution := lookup(env, key).value.(*Expr).next.next

			expr.child = substitution
			expr.child.next = listChild.next

			return eval(expr, env)
		} else if value.kind == NULL {
			// Empty
		} else {
			fmt.Println(typeof(value.kind))
			panic("evalList: unknown kind of symbol")
		}
	}

	return *expr
}

func evalAtom(expr *Expr, env *Env) Expr {
	if expr.kind == SYMBOL {
		value := lookup(env, expr.value.(string))

		if value.kind == UNKNOWN {
			fmt.Println(expr.value.(string))
			panic("eval: unknown symbol")
		}

		return value
	} else if expr.kind == NUMBER {
		return *expr
	} else if expr.kind == STRING {
		return *expr
	} else if expr.kind == PAMBDA {
		return *expr
	} else if expr.kind == BOOL {
		return *expr
	} else if expr.kind == NULL {
		return *expr
	}

	fmt.Println(typeof(expr.kind))
	panic("eval: unknown kind")
}

func eval(expr *Expr, env *Env) Expr {
	result := Expr{UNKNOWN, nil, nil, nil}

	if expr.kind == LIST {
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
