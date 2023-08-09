package core

func set(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	key := eval(&e, env)
	value := eval(e.Next, env)

	if key.Kind != SYMBOL {
		panic("set requires a symbol")
	}

	addEnv(env, key.Value.(string), eval(&value, env))

	return Expr{NULL, nil, nil, nil}
}

func define(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	key := eval(&e, env)
	value := eval(e.Next, env)

	if key.Kind != SYMBOL {
		panic("define requires a symbol")
	}

	addEnv(env, key.Value.(string), eval(&value, env))

	return Expr{NULL, nil, nil, nil}
}
