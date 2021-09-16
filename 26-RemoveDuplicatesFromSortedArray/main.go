package main

import "fmt"

func removeDuplicates(nums []int) int {
	k := 0
	curr := -101

	for _, val := range nums {
		if val != curr {
			curr = val
			nums[k] = val
			k++
		}
	}

	return k
}

func main() {
	output := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})

	fmt.Println(output)
}
