package main

import "fmt"

func isNumberPalindrome(num int) bool {
	var recur func(curr int) bool

	recur = func(curr int) bool {
		if curr >= 0 && curr < 10 {
			return curr == num%10
		}

		if !recur(curr / 10) {
			return false
		}

		num /= 10

		return curr%10 == num%10
	}

	return recur(num)
}

func main() {
	num := 32123

	output := isNumberPalindrome(num)

	fmt.Println(output)
}
