package main

import "fmt"

func subsets(nums []int) [][]int {
	count := pow(len(nums))
	sets := make([][]int, count)

	for i := 0; i < count; i++ {
		curr := []int{}

		for j, k := 0, 1; j <= i; j, k = j+1, k<<1 {
			if i&k > 0 {
				curr = append(curr, nums[j])
			}
		}

		sets[i] = curr
	}

	return sets
}

func pow(length int) int {
	count := 1

	for i := 0; i < length; i++ {
		count <<= 1
	}

	return count
}

func main() {
	output := subsets([]int{1, 2, 3, 4})

	fmt.Println(output)
}
