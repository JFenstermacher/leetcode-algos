package main

import "fmt"

func singleNumber(nums []int) int {
	mapping := map[int]int{}

	for _, num := range nums {
		_, found := mapping[num]

		if found {
			mapping[num]++
		} else {
			mapping[num] = 1
		}
	}

	for num, count := range mapping {
		if count == 1 {
			return num
		}
	}

	return -1
}

func main() {
	output := singleNumber([]int{2, 2, 1})
	fmt.Println(output)
}
