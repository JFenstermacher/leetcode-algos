package main

import "fmt"

func climbStairs(n int) int {
	values := map[int]int{
		0: 1,
		1: 1,
	}

	for i := 2; i <= n; i++ {
		values[i] = values[i-1] + values[i-2]
	}

	return values[n]
}

func main() {
	output := climbStairs(7)

	fmt.Println(output)
}
