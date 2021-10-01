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

// Totally stolen from top answer
// So smart, I'm a chump lol
func smartSingleNumber(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}

	return ans
}

func main() {
	output := smartSingleNumber([]int{2, 2, 1})
	fmt.Println(output)
}
