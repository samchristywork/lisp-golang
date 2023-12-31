package core

import (
	"fmt"
	"lisp/model"
	"strconv"
)

func findRightParen(tokens []string) *int {
	count := 1
	for i, token := range tokens {
		if token == "(" {
			count++
		} else if token == ")" {
			count--
		}
		if count == 0 {
			return &i
		}
	}

	fmt.Println("unbalanced parentheses")
	return nil
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

func handleList(expr *Expr, tokens []string) bool {
	right_paren_index := findRightParen(tokens)

	if right_paren_index == nil {
		return false
	}

	if *right_paren_index != 0 {
		child := model.EmptyListExpr()
		expr.Child = child
		buildExpressionTree(child, tokens[:*right_paren_index])
	}

	if len(tokens) > *right_paren_index+1 {
		next := model.EmptyListExpr()
		expr.Next = next
		buildExpressionTree(next, tokens[*right_paren_index+1:])
	}

	return true
}

func handleAtom(expr *Expr, token string, tokens []string) {
	if len(token) > 0 && token[0] == '"' { // Handle Strings
		expr.Kind = model.STRING
		expr.Value = token[1 : len(token)-1]

	} else if token == "true" || token == "false" { // Handle Booleans
		expr.Kind = model.BOOL
		expr.Value = token == "true"

	} else { // Handle Numbers and Symbols
		expr.Kind = model.NUMBER
		val, e := strconv.ParseFloat(token, 64)

		expr.Value = val

		if e != nil {
			expr.Kind = model.SYMBOL
			expr.Value = token
		}
	}

	if len(tokens) > 0 {
		next := model.EmptyListExpr()
		expr.Next = next
		buildExpressionTree(next, tokens)
	}
}

func buildExpressionTree(expr *Expr, tokens []string) bool {
	token := tokens[0]
	tokens = tokens[1:]

	if token == "(" {
		if !handleList(expr, tokens) {
			return false
		}
	} else if token == ")" {
		fmt.Println("unexpected )")
		return false

	} else {
		handleAtom(expr, token, tokens)
	}

	return true
}

func Parse(input string) *Expr {
	exp := model.EmptyListExpr()
	tokens := tokenize(input)
	if !buildExpressionTree(exp, tokens) {
		return nil
	}

	return exp
}
