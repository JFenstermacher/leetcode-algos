package main

import "fmt"

func atoi(s string) int {
	var recur func(n int) int

	recur = func(n int) int {
		if n == 1 {
			return int(s[n-1] - '0')
		}

		return 10*recur(n-1) + int(s[n-1]-'0')
	}

	return recur(len(s))
}

func main() {
	output := atoi("54321")

	fmt.Println(output)
}
