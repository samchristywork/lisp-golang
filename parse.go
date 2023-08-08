package main

import (
	"strconv"
)

func findRightParen(tokens []string) int {
	count := 1
	for i, token := range tokens {
		if token == "(" {
			count++
		} else if token == ")" {
			count--
		}
		if count == 0 {
			return i
		}
	}
	panic("unbalanced parentheses")
}

func tokenize(input string) []string {
	tokens := []string{}

	for n := 0; n < len(input); n++ {
		character := input[n]

		if character == '(' { // Handle Parentheses
			tokens = append(tokens, "(")
		} else if character == ')' {
			tokens = append(tokens, ")")

		} else if character == ' ' || character == '\n' { // Handle Whitespace
			continue

		} else if character == '"' { // Handle Strings
			for m := n + 1; m < len(input); m++ {
				if input[m] != '"' {
					continue
				}

				tokens = append(tokens, input[n:m+1])
				n = m
				break
			}

		} else { // Handle other Atoms
			for m := n + 1; m < len(input); m++ {
				if input[m] != '(' && input[m] != ')' && input[m] != ' ' && input[m] != '\n' {
					continue
				}

				tokens = append(tokens, input[n:m])
				n = m - 1
				break
			}
		}
	}

	return tokens
}

func readFromTokens(expr *Expr, tokens []string) {
	if len(tokens) == 0 {
		next := Expr{LIST, nil, nil, nil}
		expr.next = &next
		return
	}

	token := tokens[0]
	tokens = tokens[1:]

	if token == "(" { // Handle Lists
		right_paren_index := findRightParen(tokens)
		child := Expr{LIST, nil, nil, nil}
		expr.child = &child
		readFromTokens(&child, tokens[:right_paren_index])
		if len(tokens) > right_paren_index+1 {
			next := Expr{LIST, nil, nil, nil}
			expr.next = &next
			readFromTokens(&next, tokens[right_paren_index+1:])
		}
	} else if token == ")" {
		panic("unexpected )")

	} else { // Handle Atoms
		if len(token) > 0 && token[0] == '"' { // Handle Strings
			expr.kind = STRING
			expr.value = token[1 : len(token)-1]

		} else { // Handle Numbers and Symbols
			expr.kind = NUMBER
			val, e := strconv.ParseFloat(token, 64)

			expr.value = val

			if e != nil {
				expr.kind = SYMBOL
				expr.value = token
			}
		}
		if len(tokens) > 0 {
			next := Expr{LIST, nil, nil, nil}
			expr.next = &next
			readFromTokens(&next, tokens)
		}
	}
}

func parse(input string) *Expr {
	exp := Expr{LIST, nil, nil, nil}
	readFromTokens(&exp, tokenize(input))
	return &exp
}
