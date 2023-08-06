package main

func and(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)

	if a.kind != BOOL || b.kind != BOOL {
		panic("and requires two booleans")
	}

	return Expr{BOOL, a.value.(bool) && b.value.(bool), nil, nil}
}

func or(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)

	if a.kind != BOOL || b.kind != BOOL {
		panic("and requires two booleans")
	}

	return Expr{BOOL, a.value.(bool) || b.value.(bool), nil, nil}
}

func not(e Expr, env *Env) Expr {
	a := eval(&e, env)

	if a.kind != BOOL {
		panic("not requires a boolean")
	}

	return Expr{BOOL, !a.value.(bool), nil, nil}
}

func xor(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)

	if a.kind != BOOL || b.kind != BOOL {
		panic("xor requires two booleans")
	}

	left := a.value.(bool) && !b.value.(bool)
	right := !a.value.(bool) && b.value.(bool)

	return Expr{BOOL, left || right, nil, nil}
}

func nor(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)

	if a.kind != BOOL || b.kind != BOOL {
		panic("nor requires two booleans")
	}

	return Expr{BOOL, !(a.value.(bool) || b.value.(bool)), nil, nil}
}

func nand(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)

	if a.kind != BOOL || b.kind != BOOL {
		panic("nand requires two booleans")
	}

	return Expr{BOOL, !(a.value.(bool) && b.value.(bool)), nil, nil}
}

func xnor(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)

	if a.kind != BOOL || b.kind != BOOL {
		panic("xnor requires two booleans")
	}

	left := a.value.(bool) && !b.value.(bool)
	right := !a.value.(bool) && b.value.(bool)

	return Expr{BOOL, !(left || right), nil, nil}
}
