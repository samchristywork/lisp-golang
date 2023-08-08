package main

func equals(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)
	a.next = nil
	b.next = nil

	if a.kind != b.kind {
		return Expr{BOOL, false, nil, nil}
	}

	if a.value == b.value {
		return Expr{BOOL, true, nil, nil}
	} else {
		return Expr{BOOL, false, nil, nil}
	}
}

func notEquals(e Expr, env *Env) Expr {
	if equals(e, env).value.(bool) {
		return Expr{BOOL, false, nil, nil}
	}

	return Expr{BOOL, true, nil, nil}
}

func lessThan(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)
	a.next = nil
	b.next = nil

	if a.kind != NUMBER || b.kind != NUMBER {
		panic("lessThan requires two numbers")
	}

	return Expr{BOOL, a.value.(float64) < b.value.(float64), nil, nil}
}

func greaterThan(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)
	a.next = nil
	b.next = nil

	if a.kind != NUMBER || b.kind != NUMBER {
		panic("greaterThan requires two numbers")
	}

	return Expr{BOOL, a.value.(float64) > b.value.(float64), nil, nil}
}

func lessThanEquals(e Expr, env *Env) Expr {
	lt := lessThan(e, env)
	eq := equals(e, env)
	lt.next = nil
	eq.next = nil

	if lt.value.(bool) || eq.value.(bool) {
		return Expr{BOOL, true, nil, nil}
	} else {
		return Expr{BOOL, false, nil, nil}
	}
}

func greaterThanEquals(e Expr, env *Env) Expr {
	gr := greaterThan(e, env)
	eq := equals(e, env)
	gr.next = nil
	eq.next = nil

	if gr.value.(bool) || eq.value.(bool) {
		return Expr{BOOL, true, nil, nil}
	} else {
		return Expr{BOOL, false, nil, nil}
	}
}
