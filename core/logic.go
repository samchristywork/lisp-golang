package core

func and(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.Next, env)

	if a.Kind != BOOL || b.Kind != BOOL {
		panic("and requires two booleans")
	}

	return Expr{BOOL, a.Value.(bool) && b.Value.(bool), nil, nil}
}

func or(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.Next, env)

	if a.Kind != BOOL || b.Kind != BOOL {
		panic("and requires two booleans")
	}

	return Expr{BOOL, a.Value.(bool) || b.Value.(bool), nil, nil}
}

func not(e Expr, env *Env) Expr {
	a := eval(&e, env)

	if a.Kind != BOOL {
		panic("not requires a boolean")
	}

	return Expr{BOOL, !a.Value.(bool), nil, nil}
}

func xor(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.Next, env)

	if a.Kind != BOOL || b.Kind != BOOL {
		panic("xor requires two booleans")
	}

	left := a.Value.(bool) && !b.Value.(bool)
	right := !a.Value.(bool) && b.Value.(bool)

	return Expr{BOOL, left || right, nil, nil}
}

func nor(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.Next, env)

	if a.Kind != BOOL || b.Kind != BOOL {
		panic("nor requires two booleans")
	}

	return Expr{BOOL, !(a.Value.(bool) || b.Value.(bool)), nil, nil}
}

func nand(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.Next, env)

	if a.Kind != BOOL || b.Kind != BOOL {
		panic("nand requires two booleans")
	}

	return Expr{BOOL, !(a.Value.(bool) && b.Value.(bool)), nil, nil}
}

func xnor(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.Next, env)

	if a.Kind != BOOL || b.Kind != BOOL {
		panic("xnor requires two booleans")
	}

	left := a.Value.(bool) && !b.Value.(bool)
	right := !a.Value.(bool) && b.Value.(bool)

	return Expr{BOOL, !(left || right), nil, nil}
}
