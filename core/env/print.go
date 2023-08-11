package env

import (
	"fmt"
	"sort"
)

func PrintNode(expr *Expr) {
	fmt.Printf("%s ", typeof(expr.Kind))
	fmt.Printf(" %v ", expr.Value)
	if expr.Next != nil {
		fmt.Printf("next: %v ", expr.Next)
	}
	if expr.Child != nil {
		fmt.Printf("child: %v", expr.Child)
	}
	fmt.Println()
}

func _printExpr(expr *Expr, depth int) {
	if expr == nil {
		return
	}

	fmt.Print(typeof(expr.Kind))

	for i := 0; i < 7-len(typeof(expr.Kind)); i++ {
		fmt.Print(" ")
	}
	fmt.Print("| ")

	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}

	if expr.Kind == LIST {
		fmt.Println("(")
	} else {
		fmt.Println(expr.Value)
	}

	if expr.Child != nil {
		_printExpr(expr.Child, depth+1)
	}

	if expr.Next != nil {
		_printExpr(expr.Next, depth)
	}
}

func PrintExpr(expr *Expr) {
	PrintNode(expr)

	if expr.Child != nil {
		_printExpr(expr.Child, 1)
	}
	fmt.Println()
}

func _printEnv(env *Env, depth int) {
	if env.Outer != nil {
		_printEnv(env.Outer, depth+1)
	}

	keys := make([]string, 0, len(env.Data))

	for key := range env.Data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%d\t", depth)
		fmt.Printf("%s\t", key)
		PrintNode(env.Data[key])
	}
}

func PrintEnv(env *Env) {
	fmt.Println("Scope\tLabel\tType      Value")
	_printEnv(env, 0)
}
