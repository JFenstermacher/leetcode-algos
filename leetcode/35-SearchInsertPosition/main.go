package main

import "fmt"

func searchInsert(nums []int, target int) int {
	min := 0
	max := len(nums) - 1
	guess := int((min + max) / 2)

	for min < max {
		curr := nums[guess]

		if curr == target {
			return guess
		} else if curr < target {
			min = guess + 1
		} else {
			max = guess - 1
		}

		guess = int((min + max) / 2)
	}

	if nums[guess] < target {
		return guess + 1
	} else {
		return guess
	}
}

func main() {
	input := []int{1, 3, 5, 6}
	target := 0

	output := searchInsert(input, target)

	fmt.Println(output)
}
