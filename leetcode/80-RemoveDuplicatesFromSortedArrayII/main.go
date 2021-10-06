package main

import "fmt"

func removeDuplicates(nums []int) int {
	k, index, curr, count := 0, 0, nums[0], 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == curr {
			count++
		} else {
			curr = nums[i]
			count = 1
		}

		if count > 2 {
			continue
		}

		nums[index] = nums[i]
		index++
		k++
	}

	return k
}

// Stolen from top answer
func clever(nums []int) int {
	cur := 0

	for i, num := range nums {
		if i < 2 || num > nums[cur-2] {
			nums[cur] = num
			cur++
		}
	}

	return cur
}

func main() {
	input := []int{1, 1, 1, 2, 2, 3}

	output := removeDuplicates(input)

	fmt.Println(output, input)
}
