package main

import (
	"os"
)

func main() {
	if len(os.Args) > 1 {
		file()

		return
	}

	repl()
}
