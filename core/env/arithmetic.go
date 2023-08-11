package env

func plus(operands []*Expr, env *Env, evaluator Callback) *Expr {
	if len(operands) != 2 {
		panic("plus requires two operands")
	}

	a := evaluator(operands[0], env)
	b := evaluator(operands[1], env)

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("plus requires two numbers")
	}

	return &Expr{Kind: NUMBER, Value: a.Value.(float64) + b.Value.(float64), Next: nil, Child: nil}
}

func minus(e *Expr, env *Env, evaluator Callback) *Expr {
	a := evaluator(e, env)
	b := evaluator(e.Next, env)
	a.Next = nil
	b.Next = nil

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("minus requires two numbers")
	}

	return &Expr{Kind: NUMBER, Value: a.Value.(float64) - b.Value.(float64), Next: nil, Child: nil}
}

func multiply(e *Expr, env *Env, evaluator Callback) *Expr {
	a := evaluator(e, env)
	b := evaluator(e.Next, env)
	a.Next = nil
	b.Next = nil

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("multiply requires two numbers")
	}

	return &Expr{Kind: NUMBER, Value: a.Value.(float64) * b.Value.(float64), Next: nil, Child: nil}
}

func divide(e *Expr, env *Env, evaluator Callback) *Expr {
	a := evaluator(e, env)
	b := evaluator(e.Next, env)
	a.Next = nil
	b.Next = nil

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("divide requires two numbers")
	}

	return &Expr{Kind: NUMBER, Value: a.Value.(float64) / b.Value.(float64), Next: nil, Child: nil}
}
