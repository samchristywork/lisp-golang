package env

func equals(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	a := evaluator(&e, env)
	b := evaluator(e.Next, env)
	a.Next = nil
	b.Next = nil

	if a.Kind != b.Kind {
		return Expr{Kind: BOOL, Value: false, Next: nil, Child: nil}
	}

	if a.Value == b.Value {
		return Expr{Kind: BOOL, Value: true, Next: nil, Child: nil}
	} else {
		return Expr{Kind: BOOL, Value: false, Next: nil, Child: nil}
	}
}

func notEquals(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	if equals(e, env, evaluator).Value.(bool) {
		return Expr{Kind: BOOL, Value: false, Next: nil, Child: nil}
	}

	return Expr{Kind: BOOL, Value: true, Next: nil, Child: nil}
}

func lessThan(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	a := evaluator(&e, env)
	b := evaluator(e.Next, env)
	a.Next = nil
	b.Next = nil

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("lessThan requires two numbers")
	}

	return Expr{Kind: BOOL, Value: a.Value.(float64) < b.Value.(float64), Next: nil, Child: nil}
}

func greaterThan(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	a := evaluator(&e, env)
	b := evaluator(e.Next, env)
	a.Next = nil
	b.Next = nil

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("greaterThan requires two numbers")
	}

	return Expr{Kind: BOOL, Value: a.Value.(float64) > b.Value.(float64), Next: nil, Child: nil}
}

func lessThanEquals(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	lt := lessThan(e, env, evaluator)
	eq := equals(e, env, evaluator)
	lt.Next = nil
	eq.Next = nil

	if lt.Value.(bool) || eq.Value.(bool) {
		return Expr{Kind: BOOL, Value: true, Next: nil, Child: nil}
	} else {
		return Expr{Kind: BOOL, Value: false, Next: nil, Child: nil}
	}
}

func greaterThanEquals(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	gr := greaterThan(e, env, evaluator)
	eq := equals(e, env, evaluator)
	gr.Next = nil
	eq.Next = nil

	if gr.Value.(bool) || eq.Value.(bool) {
		return Expr{Kind: BOOL, Value: true, Next: nil, Child: nil}
	} else {
		return Expr{Kind: BOOL, Value: false, Next: nil, Child: nil}
	}
}
