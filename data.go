package main

func set(e Expr, env *Env) Expr {
	key := eval(&e, env)
	value := eval(e.next, env)

	if key.kind != SYMBOL {
		panic("set requires a symbol")
	}

	addEnv(env, key.value.(string), eval(&value, env))

	return Expr{NULL, nil, nil, nil}
}

func define(e Expr, env *Env) Expr {
	key := eval(&e, env)
	value := eval(e.next, env)

	if key.kind != SYMBOL {
		panic("define requires a symbol")
	}

	addEnv(env, key.value.(string), eval(&value, env))

	return Expr{NULL, nil, nil, nil}
}
