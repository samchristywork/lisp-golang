package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	env := initEnv()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		if input == "env" || input == "." {
			printEnv(env)
			continue
		}

		if input == "exit" {
			break
		}

		if len(input) == 0 {
			break
		}

		expression := parse(input)
		value := eval(expression, env)
		fmt.Println(value.value)
	}
}

func file() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	content := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	env := initEnv()
	expression := parse("(begin \n" + content + ")")

	value := eval(expression, env)
	fmt.Println(value.value)
}
