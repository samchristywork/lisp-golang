package env

func equals(operands []*Expr, env *Env, evaluator Callback) *Expr {
	if len(operands) != 2 {
		panic("equals requires two operands")
	}

	a := evaluator(operands[0], env)
	b := evaluator(operands[1], env)

	if a.Kind != b.Kind {
		return &Expr{Kind: BOOL, Value: false, Next: nil, Child: nil}
	}

	if a.Value == b.Value {
		return &Expr{Kind: BOOL, Value: true, Next: nil, Child: nil}
	} else {
		return &Expr{Kind: BOOL, Value: false, Next: nil, Child: nil}
	}
}

func notEquals(operands []*Expr, env *Env, evaluator Callback) *Expr {
	if equals(operands, env, evaluator).Value.(bool) {
		return &Expr{Kind: BOOL, Value: false, Next: nil, Child: nil}
	}

	return &Expr{Kind: BOOL, Value: true, Next: nil, Child: nil}
}

func lessThan(e *Expr, env *Env, evaluator Callback) *Expr {
	a := evaluator(e, env)
	b := evaluator(e.Next, env)
	a.Next = nil
	b.Next = nil

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("lessThan requires two numbers")
	}

	return &Expr{Kind: BOOL, Value: a.Value.(float64) < b.Value.(float64), Next: nil, Child: nil}
}

func greaterThan(e *Expr, env *Env, evaluator Callback) *Expr {
	a := evaluator(e, env)
	b := evaluator(e.Next, env)
	a.Next = nil
	b.Next = nil

	if a.Kind != NUMBER || b.Kind != NUMBER {
		panic("greaterThan requires two numbers")
	}

	return &Expr{Kind: BOOL, Value: a.Value.(float64) > b.Value.(float64), Next: nil, Child: nil}
}

//func lessThanEquals(e *Expr, env *Env, evaluator Callback) *Expr {
//	a := evaluator(e, env)
//
//	if lt.Value.(bool) || eq.Value.(bool) {
//		return &Expr{Kind: BOOL, Value: true, Next: nil, Child: nil}
//	} else {
//		return &Expr{Kind: BOOL, Value: false, Next: nil, Child: nil}
//	}
//}
//
//func greaterThanEquals(e *Expr, env *Env, evaluator Callback) *Expr {
//	gr := greaterThan(e, env, evaluator)
//	eq := equals(e, env, evaluator)
//	gr.Next = nil
//	eq.Next = nil
//
//	if gr.Value.(bool) || eq.Value.(bool) {
//		return &Expr{Kind: BOOL, Value: true, Next: nil, Child: nil}
//	} else {
//		return &Expr{Kind: BOOL, Value: false, Next: nil, Child: nil}
//	}
//}
