package core

func equals(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	a := eval(&e, env)
	b := eval(e.Next, env)
	a.Next = nil
	b.Next = nil

	if a.Kind != b.Kind {
		return Expr{BOOL, false, nil, nil}
	}

	if a.Value == b.Value {
		return Expr{BOOL, true, nil, nil}
	} else {
		return Expr{BOOL, false, nil, nil}
	}
}

func notEquals(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	if equals(e, env, evaluator).Value.(bool) {
		return Expr{BOOL, false, nil, nil}
	}

	return Expr{BOOL, true, nil, nil}
}

func lessThan(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	a := eval(&e, env)
	b := eval(e.Next, env)
	a.Next = nil
	b.Next = nil

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("lessThan requires two numbers")
	}

	return Expr{BOOL, a.Value.(float64) < b.Value.(float64), nil, nil}
}

func greaterThan(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	a := eval(&e, env)
	b := eval(e.Next, env)
	a.Next = nil
	b.Next = nil

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("greaterThan requires two numbers")
	}

	return Expr{BOOL, a.Value.(float64) > b.Value.(float64), nil, nil}
}

func lessThanEquals(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	lt := lessThan(e, env, evaluator)
	eq := equals(e, env, evaluator)
	lt.Next = nil
	eq.Next = nil

	if lt.Value.(bool) || eq.Value.(bool) {
		return Expr{BOOL, true, nil, nil}
	} else {
		return Expr{BOOL, false, nil, nil}
	}
}

func greaterThanEquals(e Expr, env *Env, evaluator func(*Expr, *Env) Expr) Expr {
	gr := greaterThan(e, env, evaluator)
	eq := equals(e, env, evaluator)
	gr.Next = nil
	eq.Next = nil

	if gr.Value.(bool) || eq.Value.(bool) {
		return Expr{BOOL, true, nil, nil}
	} else {
		return Expr{BOOL, false, nil, nil}
	}
}
