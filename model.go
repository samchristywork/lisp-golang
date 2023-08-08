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

func (e *Expr) String() string {
	return types[e.kind]
}

func (e *Expr) Equal(other *Expr) bool {
	if e.kind != other.kind {
		return false
	}

	switch e.kind {
	case NULL:
		return e.value == other.value
	case BOOL:
		return e.value.(bool) == other.value.(bool)
	case NUMBER:
		return e.value.(float64) == other.value.(float64)
	case STRING:
		return e.value.(string) == other.value.(string)
	case SYMBOL:
		return e.value.(string) == other.value.(string)
	}

	if e.child != nil || other.child != nil {
		return e.child.Equal(other.child)
	}

	if e.next != nil || other.next != nil {
		return e.next.Equal(other.next)
	}

	return false
}
