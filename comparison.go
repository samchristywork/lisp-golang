package main

func equals(e Expr, env *Env) Expr {
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

func notEquals(e Expr, env *Env) Expr {
	if equals(e, env).Value.(bool) {
		return Expr{BOOL, false, nil, nil}
	}

	return Expr{BOOL, true, nil, nil}
}

func lessThan(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.Next, env)
	a.Next = nil
	b.Next = nil

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("lessThan requires two numbers")
	}

	return Expr{BOOL, a.Value.(float64) < b.Value.(float64), nil, nil}
}

func greaterThan(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.Next, env)
	a.Next = nil
	b.Next = nil

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("greaterThan requires two numbers")
	}

	return Expr{BOOL, a.Value.(float64) > b.Value.(float64), nil, nil}
}

func lessThanEquals(e Expr, env *Env) Expr {
	lt := lessThan(e, env)
	eq := equals(e, env)
	lt.Next = nil
	eq.Next = nil

	if lt.Value.(bool) || eq.Value.(bool) {
		return Expr{BOOL, true, nil, nil}
	} else {
		return Expr{BOOL, false, nil, nil}
	}
}

func greaterThanEquals(e Expr, env *Env) Expr {
	gr := greaterThan(e, env)
	eq := equals(e, env)
	gr.Next = nil
	eq.Next = nil

	if gr.Value.(bool) || eq.Value.(bool) {
		return Expr{BOOL, true, nil, nil}
	} else {
		return Expr{BOOL, false, nil, nil}
	}
}
