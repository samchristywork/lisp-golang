package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Golang Lisp Interpreter")

	if len(os.Args) > 1 {
		file()

		return
	}

	repl()
}
