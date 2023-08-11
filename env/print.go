package env

import (
	"fmt"
	"lisp/model"
	"lisp/util"
	"sort"
)

// Similar to PrintExpr, but does not recurse
func PrintNode(expr *Expr) {
	// TODO: Implement
}

func _printExpr(expr *Expr) {
	if expr == nil {
		return
	}

	if expr.Kind == model.NUMBER {
		util.Blue()
		fmt.Printf("%v", expr.Value)
		util.Reset()
	} else if expr.Kind == model.STRING {
		fmt.Printf("\"")
		util.Green()
		fmt.Printf("%v", expr.Value)
		util.Reset()
		fmt.Printf("\"")
	} else if expr.Kind == model.BOOL {
		util.Magenta()
		fmt.Printf("%v", expr.Value)
		util.Reset()
	} else if expr.Kind == model.SYMBOL {
		util.Cyan()
		fmt.Printf("%v", expr.Value)
		util.Reset()
	} else if expr.Kind == model.NULL {
		util.Red()
		fmt.Printf("%v", expr.Value)
		util.Reset()
	} else if expr.Kind == model.LIST {
		fmt.Printf("(")
		_printExpr(expr.Child)
		fmt.Printf(")")
	} else if expr.Kind == model.FUNCTION {
		util.Yellow()
		fmt.Printf("%v", expr.Child)
		util.Reset()
	} else if expr.Kind == model.LAMBDA {
		util.Blue()
		fmt.Printf("%v", expr.Value)
		util.Reset()
	}

	if expr.Next != nil {
		fmt.Printf(" ")
		_printExpr(expr.Next)
	}
}

func PrintExpr(expr *Expr) {
	if expr != nil {
		_printExpr(expr)
	}
	fmt.Println()
}

func PrintEnvLine(expr *Expr) {
	fmt.Printf("%s ", util.Typeof(expr.Kind))
	fmt.Printf(" %v ", expr.Value)
	if expr.Next != nil {
		fmt.Printf("next: %v ", expr.Next)
	}
	if expr.Child != nil {
		fmt.Printf("child: %v", expr.Child)
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
		PrintEnvLine(env.Data[key])
	}
}

func PrintEnv(env *Env) {
	fmt.Println("Scope\tLabel\tType      Value")
	_printEnv(env, 0)
}
