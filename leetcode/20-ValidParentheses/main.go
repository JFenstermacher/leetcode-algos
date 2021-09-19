package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		last := stack[len(stack)-1]

		if (c == ')' && last != '(') || (c == ']' && last != '[') || (c == '}' && last != '{') {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func main() {
	output := isValid("{[]}")

	if output {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}
