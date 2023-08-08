package main

const (
	UNKNOWN = iota
	BOOL
	FUNCTION
	LAMBDA
	LIST
	NULL
	NUMBER
	PAMBDA
	STRING
	SYMBOL
)

var types = map[int]string{
	UNKNOWN:  "Unknown ",
	BOOL:     "Bool    ",
	FUNCTION: "Function",
	LAMBDA:   "Lambda  ",
	LIST:     "List    ",
	NULL:     "Null    ",
	NUMBER:   "Number  ",
	PAMBDA:   "Pambda  ",
	STRING:   "String  ",
	SYMBOL:   "Symbol  ",
}

type Expr struct {
	kind  int
	value interface{}
	next  *Expr
	child *Expr
}
