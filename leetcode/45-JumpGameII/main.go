package main

import (
	"fmt"
	"math"
)

func jump(nums []int) int {
	res := make([]int, len(nums))

	for i := 0; i < len(res); i++ {
		res[i] = math.MaxInt64
	}

	res[0] = 0
	for i, num := range nums {
		right := min(i+num, len(nums)-1)

		for j := right; j > i; j-- {
			res[j] = min(res[i]+1, res[j])
		}
	}

	return res[len(res)-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	output := jump([]int{2, 3, 0, 1, 4})

	fmt.Println(output)
}
