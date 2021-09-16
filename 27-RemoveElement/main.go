package main

import "fmt"

func removeElement(nums []int, val int) int {
	index := 0

	for _, num := range nums {
		if num != val {
			nums[index] = num
			index++
		}
	}

	return index
}

func main() {
	output := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)

	fmt.Println(output)
}
