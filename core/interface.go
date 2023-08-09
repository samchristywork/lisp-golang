package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Repl() {
	env := initEnv()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		if input == "help" {
			fmt.Println("help: print this help message")
			fmt.Println("env: print the environment")
			fmt.Println("exit: exit the repl")
			continue
		}

		if len(input) != 0 {
			input = strings.Split(input, ";")[0]
		}

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
		fmt.Println(value.Value)
	}
}

func File() {
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
		line := scanner.Text()
		if len(line) != 0 {
			line = strings.Split(line, ";")[0]
		}
		content += line + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	env := initEnv()
	expression := parse("(begin (print \"Program Begin\")\n" + content + ")")

	eval(expression, env)
}
