package env

func expectTwoNumbers(operands []*Expr, env *Env, evaluator Callback) (*Expr, *Expr) {
	if len(operands) != 2 {
		panic("requires two operands")
	}

	a := evaluator(operands[0], env)
	b := evaluator(operands[1], env)

	if operands[0].Kind != NUMBER || operands[1].Kind != NUMBER {
		panic("requires two numbers")
	}

	return a, b
}

func expectTwoOperands(operands []*Expr, env *Env, evaluator Callback) (*Expr, *Expr) {
	if len(operands) != 2 {
		panic("requires two operands")
	}

	a := evaluator(operands[0], env)
	b := evaluator(operands[1], env)

	return a, b
}
