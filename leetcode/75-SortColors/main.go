package main

import "fmt"

func sortColors(nums []int) {
	low, high := 0, len(nums)-1

	// Zeros in front
	for low < high {
		for ; low < high; low++ {
			if nums[low] != 0 {
				break
			}
		}

		if nums[high] == 0 {
			nums[low], nums[high] = nums[high], nums[low]
		}

		high--
	}

	high = len(nums) - 1

	// 2s in back
	for low < high {
		for ; low < high; high-- {
			if nums[high] != 2 {
				break
			}
		}

		if nums[low] == 2 {
			nums[low], nums[high] = nums[high], nums[low]
		}

		low++
	}
}

func main() {
	nums := []int{1}

	sortColors(nums)

	fmt.Println(nums)
}
