package model

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

var Types = map[int]string{
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
	Kind  int
	Value interface{}
	Next  *Expr
	Child *Expr
}

func (e *Expr) String() string {
	return Types[e.Kind]
}

func (e *Expr) Equal(other *Expr) bool {
	if e.Kind != other.Kind {
		return false
	}

	switch e.Kind {
	case NULL:
		return e.Value == other.Value
	case BOOL:
		return e.Value.(bool) == other.Value.(bool)
	case NUMBER:
		return e.Value.(float64) == other.Value.(float64)
	case STRING:
		return e.Value.(string) == other.Value.(string)
	case SYMBOL:
		return e.Value.(string) == other.Value.(string)
	}

	if e.Child != nil || other.Child != nil {
		return e.Child.Equal(other.Child)
	}

	if e.Next != nil || other.Next != nil {
		return e.Next.Equal(other.Next)
	}

	return false
}
