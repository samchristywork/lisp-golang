package main

import (
	"fmt"
	"lisp/core"
	"os"
)

func main() {
	fmt.Println("Golang Lisp Interpreter")

	if len(os.Args) > 1 {
		core.File()

		return
	}

	core.Repl()
}
