package env

import ()

func set(e *Expr, env *Env, evaluator Callback) *Expr {
	key := evaluator(e, env)
	value := evaluator(e.Next, env)

	if key.Kind != SYMBOL {
		panic("set requires a symbol")
	}

	AddEnv(env, key.Value.(string), evaluator(value, env))

	return &Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}
}

func define(operands []*Expr, env *Env, evaluator Callback) *Expr {
	if len(operands) != 2 {
		panic("define requires 2 operands")
	}

	key := operands[0]
	value := operands[1]

	if key.Kind == SYMBOL {
		AddEnv(env, key.Value.(string), value)
	} else if key.Kind == LIST {
		foo := &Expr{Kind: SHARK, Value: key, Next: nil, Child: value}
		AddEnv(env, key.Child.Value.(string), foo)
	} else {
		panic("define requires a symbol")
	}

	//if e.Kind == SYMBOL {
	//	key := e.Value.(string)
	//	value := evaluator(e.Next, env)

	//	AddEnv(env, key, value)
	//} else if e.Kind == LIST {
	//	key := e.Child.Value.(string)

	//	AddEnv(env, key, e.Next.Child)
	//} else {
	//	panic("define requires a symbol")
	//}

	return &Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}
}
