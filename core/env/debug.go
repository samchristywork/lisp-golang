package env

import (
	"fmt"
)

func showEnv(operands []*Expr, env *Env, evaluator Callback) *Expr {
	if len(operands) == 0 {
		PrintEnv(env)

		return &Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}
	}

	if len(operands) != 1 {
		panic("env requires one argument")
	}

	key := operands[0]

	if key.Kind != SYMBOL {
		panic("env requires a symbol")
	}

	value := Lookup(env, key.Value.(string))

	fmt.Printf("%s: ", key.Value.(string))
	PrintNode(value)

	return &Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}
}

func inspect(e *Expr, env *Env, evaluator Callback) *Expr {
	// TODO: Fix
	//e = evaluator(&e, env)

	//fmt.Printf("Expr: %v\n", e)
	//fmt.Printf("Kind: %v\n", typeof(e.Kind))
	//fmt.Printf("Value: %v\n", e.Value)
	//fmt.Printf("Child: %v\n", e.Child)
	//fmt.Printf("Next: %v\n", e.Next)

	//printExpr(e.Value.(*Expr))

	return &Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}
}

func lookupValue(e *Expr, env *Env, evaluator Callback) *Expr {
	e = evaluator(e, env)
	value := Lookup(env, e.Value.(string))
	return value
}
