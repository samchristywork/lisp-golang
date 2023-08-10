package core

import (
	environment "lisp/core/env"
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

func handleList(expr *Expr, tokens []string) {
	right_paren_index := findRightParen(tokens)

	if right_paren_index != 0 {
		child := environment.Expr{Kind: environment.LIST, Value: nil, Next: nil, Child: nil}
		expr.Child = &child
		buildExpressionTree(&child, tokens[:right_paren_index])
	}

	if len(tokens) > right_paren_index+1 {
		next := Expr{Kind: environment.LIST, Value: nil, Next: nil, Child: nil}
		expr.Next = &next
		buildExpressionTree(&next, tokens[right_paren_index+1:])
	}
}

func handleAtom(expr *Expr, token string, tokens []string) {
	if len(token) > 0 && token[0] == '"' { // Handle Strings
		expr.Kind = environment.STRING
		expr.Value = token[1 : len(token)-1]

	} else { // Handle Numbers and Symbols
		expr.Kind = environment.NUMBER
		val, e := strconv.ParseFloat(token, 64)

		expr.Value = val

		if e != nil {
			expr.Kind = environment.SYMBOL
			expr.Value = token
		}
	}
	if len(tokens) > 0 {
		next := Expr{Kind: environment.LIST, Value: nil, Next: nil, Child: nil}
		expr.Next = &next
		buildExpressionTree(&next, tokens)
	}
}

func buildExpressionTree(expr *Expr, tokens []string) {
	token := tokens[0]
	tokens = tokens[1:]

	if token == "(" {
		handleList(expr, tokens)
	} else if token == ")" {
		panic("unexpected )")

	} else {
		handleAtom(expr, token, tokens)
	}
}

func parse(input string) *Expr {
	exp := environment.Expr{Kind: environment.LIST, Value: nil, Next: nil, Child: nil}
	tokens := tokenize(input)
	buildExpressionTree(&exp, tokens)
	return &exp
}
