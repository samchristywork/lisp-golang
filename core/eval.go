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

	if operator.Kind == model.FUNCTION {
		f := operator.Value.(func([]*Expr, *environment.Env, environment.Callback) *Expr)

		return f(operands, env, eval)
	} else if operator.Kind == model.LAMBDA {
		operand := operator.Value.(*Expr).Child.Next
		body := operator.Child

		env := environment.NewEnv(env)
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

	if listChild.Kind == model.SYMBOL {
		return apply(expr.Child, env)
	}

	return expr
}

func evalAtom(expr *Expr, env *environment.Env) *Expr {
	if expr.Kind == model.SYMBOL {
		value := environment.Lookup(env, expr.Value.(string))

		if value.Kind == model.UNKNOWN {
			panic("eval: unknown symbol")
		}

		if value.Kind == model.SYMBOL && value.Value.(string) == expr.Value.(string) {
			panic("eval: loop detected")
		}

		return eval(value, env)
	} else if expr.Kind == model.NUMBER {
		return expr
	} else if expr.Kind == model.STRING {
		return expr
	} else if expr.Kind == model.BOOL {
		return expr
	} else if expr.Kind == model.NULL {
		return expr
	}

	panic("eval: unknown kind")
}

func eval(expr *Expr, env *environment.Env) *Expr {
	if expr == nil {
		return expr
	}

	if expr.Kind == model.LIST {
		return evalList(expr, env)
	} else {
		return evalAtom(expr, env)
	}
}

func EvalNew(expr *Expr) *Expr {
	env := environment.InitEnv()

	return eval(expr, env)
}
