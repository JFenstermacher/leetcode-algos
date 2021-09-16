package main

import "fmt"

func generateParentheses(n int) []string {
	result := []string{}

	var backtrack func(arr []rune, open int, closed int)

	backtrack = func(arr []rune, open int, closed int) {
		if len(arr) == 2*n {
			result = append(result, string(arr))
		}

		if open < n {
			arr = append(arr, '(')
			backtrack(arr, open+1, closed)
			arr = arr[:len(arr)-1]
		}

		if closed < open {
			arr = append(arr, ')')
			backtrack(arr, open, closed+1)
			arr = arr[:len(arr)-1]
		}
	}

	backtrack([]rune{}, 0, 0)

	return result
}

func main() {
	output := generateParentheses(2)

	fmt.Println(output)
}
