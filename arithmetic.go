package main

func plus(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)
	a.next = nil
	b.next = nil

	if a.kind != NUMBER || b.kind != NUMBER {
		panic("plus requires two numbers")
	}

	return Expr{NUMBER, a.value.(float64) + b.value.(float64), nil, nil}
}

func minus(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)
	a.next = nil
	b.next = nil

	if a.kind != NUMBER || b.kind != NUMBER {
		panic("minus requires two numbers")
	}

	return Expr{NUMBER, a.value.(float64) - b.value.(float64), nil, nil}
}

func multiply(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)
	a.next = nil
	b.next = nil

	if a.kind != NUMBER || b.kind != NUMBER {
		panic("multiply requires two numbers")
	}

	return Expr{NUMBER, a.value.(float64) * b.value.(float64), nil, nil}
}

func divide(e Expr, env *Env) Expr {
	a := eval(&e, env)
	b := eval(e.next, env)
	a.next = nil
	b.next = nil

	if a.kind != NUMBER || b.kind != NUMBER {
		panic("divide requires two numbers")
	}

	return Expr{NUMBER, a.value.(float64) / b.value.(float64), nil, nil}
}
