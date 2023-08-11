package env

func and(e *Expr, env *Env, evaluator Callback) *Expr {
	a := evaluator(e, env)
	b := evaluator(e.Next, env)

	if a.Kind != BOOL || b.Kind != BOOL {
		panic("and requires two booleans")
	}

	return &Expr{Kind: BOOL, Value: a.Value.(bool) && b.Value.(bool), Next: nil, Child: nil}
}

func or(e *Expr, env *Env, evaluator Callback) *Expr {
	a := evaluator(e, env)
	b := evaluator(e.Next, env)

	if a.Kind != BOOL || b.Kind != BOOL {
		panic("and requires two booleans")
	}

	return &Expr{Kind: BOOL, Value: a.Value.(bool) || b.Value.(bool), Next: nil, Child: nil}
}

func not(e *Expr, env *Env, evaluator Callback) *Expr {
	a := evaluator(e, env)

	if a.Kind != BOOL {
		panic("not requires a boolean")
	}

	return &Expr{Kind: BOOL, Value: !a.Value.(bool), Next: nil, Child: nil}
}

func xor(e *Expr, env *Env, evaluator Callback) *Expr {
	a := evaluator(e, env)
	b := evaluator(e.Next, env)

	if a.Kind != BOOL || b.Kind != BOOL {
		panic("xor requires two booleans")
	}

	left := a.Value.(bool) && !b.Value.(bool)
	right := !a.Value.(bool) && b.Value.(bool)

	return &Expr{Kind: BOOL, Value: left || right, Next: nil, Child: nil}
}

func nor(e *Expr, env *Env, evaluator Callback) *Expr {
	a := evaluator(e, env)
	b := evaluator(e.Next, env)

	if a.Kind != BOOL || b.Kind != BOOL {
		panic("nor requires two booleans")
	}

	return &Expr{Kind: BOOL, Value: !(a.Value.(bool) || b.Value.(bool)), Next: nil, Child: nil}
}

func nand(e *Expr, env *Env, evaluator Callback) *Expr {
	a := evaluator(e, env)
	b := evaluator(e.Next, env)

	if a.Kind != BOOL || b.Kind != BOOL {
		panic("nand requires two booleans")
	}

	return &Expr{Kind: BOOL, Value: !(a.Value.(bool) && b.Value.(bool)), Next: nil, Child: nil}
}

func xnor(e *Expr, env *Env, evaluator Callback) *Expr {
	a := evaluator(e, env)
	b := evaluator(e.Next, env)

	if a.Kind != BOOL || b.Kind != BOOL {
		panic("xnor requires two booleans")
	}

	left := a.Value.(bool) && !b.Value.(bool)
	right := !a.Value.(bool) && b.Value.(bool)

	return &Expr{Kind: BOOL, Value: !(left || right), Next: nil, Child: nil}
}
