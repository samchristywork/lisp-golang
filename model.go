package main

const (
	UNKNOWN = iota
	BOOL
	FUNCTION
	LIST
	NULL
	NUMBER
	STRING
	SYMBOL
)

var types = map[int]string{
	UNKNOWN:  "UNKNOWN",
	BOOL:     "BOOL",
	FUNCTION: "FUNCTION",
	LIST:     "LIST",
	NULL:     "NULL",
	NUMBER:   "NUMBER",
	STRING:   "STRING",
	SYMBOL:   "SYMBOL",
}

type Expr struct {
	kind  int
	value interface{}
	next  *Expr
	child *Expr
}
