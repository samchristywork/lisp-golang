package core

import (
	environment "lisp/env"
	"lisp/model"
)

type Expr = model.Expr

func apply(expr *Expr, env *environment.Env) *Expr {
	operator := environment.Lookup(env, expr.Value.(string))
	operand := expr.Next
	operands := []*Expr{}

	for operand != nil {
		operands = append(operands, operand)

		operand = operand.Next
	}

	if operator.Kind == environment.FUNCTION {
		f := operator.Value.(func([]*Expr, *environment.Env, environment.Callback) *Expr)

		return f(operands, env, eval)
	} else if operator.Kind == environment.SHARK {
		operand := operator.Value.(*Expr).Child.Next
		body := operator.Child

		for operand != nil {
			environment.AddEnv(env, operand.Value.(string), expr.Next)

			operand = operand.Next
			expr = expr.Next
		}

		ret := evalList(body, env)
		return ret
	} else {
		panic("apply: unknown operator")
	}
}

func evalList(expr *Expr, env *environment.Env) *Expr {
	listChild := expr.Child

	// Empty list
	if listChild == nil {
		return expr
	}

	if listChild.Kind == environment.SYMBOL {
		return apply(expr.Child, env)
	}

	return expr
}

func evalAtom(expr *Expr, env *environment.Env) *Expr {
	if expr.Kind == environment.SYMBOL {
		value := environment.Lookup(env, expr.Value.(string))

		if value.Kind == environment.UNKNOWN {
			panic("eval: unknown symbol")
		}

		return eval(value, env)
	} else if expr.Kind == environment.NUMBER {
		return expr
	} else if expr.Kind == environment.STRING {
		return expr
	} else if expr.Kind == environment.BOOL {
		return expr
	} else if expr.Kind == environment.NULL {
		return expr
	}

	panic("eval: unknown kind")
}

func eval(expr *Expr, env *environment.Env) *Expr {
	if expr == nil {
		return expr
	}

	if expr.Kind == environment.LIST {
		return evalList(expr, env)
	} else {
		return evalAtom(expr, env)
	}
}

func EvalNew(expr *Expr) *Expr {
	env := environment.InitEnv()

	return eval(expr, env)
}
