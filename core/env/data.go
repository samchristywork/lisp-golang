package core

func set(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	key := evaluator(&e, env)
	value := evaluator(e.Next, env)

	if key.Kind != SYMBOL {
		panic("set requires a symbol")
	}

	AddEnv(env, key.Value.(string), evaluator(&value, env))

	return Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}
}

func define(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	key := evaluator(&e, env)
	value := evaluator(e.Next, env)

	if key.Kind != SYMBOL {
		panic("define requires a symbol")
	}

	AddEnv(env, key.Value.(string), evaluator(&value, env))

	return Expr{Kind: NULL, Value: nil, Next: nil, Child: nil}
}
