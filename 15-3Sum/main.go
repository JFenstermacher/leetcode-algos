package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	threeSums := [][]int{}
	set := map[string]interface{}{}

	for i := 1; i < len(nums)-1; i++ {
		potential := fixedTwoSum(nums[i-1], nums[i:])

		for _, sum := range potential {
			s := fmt.Sprintf("%d %d %d", sum[0], sum[1], sum[2])

			_, found := set[s]

			if !found {
				threeSums = append(threeSums, sum)
			}

			set[s] = nil
		}
	}

	return threeSums
}

func fixedTwoSum(fixed int, nums []int) [][]int {
	mapping := map[int]int{}

	threeSums := [][]int{}
	for _, num := range nums {
		other, found := mapping[num]

		if found {
			arr := []int{fixed, other, num}
			sort.Ints(arr)
			threeSums = append(threeSums, arr)
		}

		mapping[0-fixed-num] = num
	}

	return threeSums
}

func main() {
	output := threeSum([]int{-1, 0, 1, 2, -1, -4})

	for _, o := range output {
		fmt.Printf("%d %d %d\n", o[0], o[1], o[2])
	}
}
