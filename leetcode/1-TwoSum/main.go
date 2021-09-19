package main

func TwoSum(nums []int, target int) []int {
	mapping := map[int]int{}

	for i, el := range nums {
		composite := target - el
		index, found := mapping[composite]

		if found == true {
			return []int{index, i}
		} else {
			mapping[el] = i
		}
	}

	return []int{0, 0}
}
