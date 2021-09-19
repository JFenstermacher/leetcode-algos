package main

func canJump(nums []int) bool {
	last := len(nums) - 1
	memo := map[int]bool{}

	var recur func(index int) bool

	recur = func(index int) bool {
		if index == last || nums[index] >= last-index {
			return true
		} else if index > last {
			return false
		}

		_, found := memo[index]

		if found {
			return false
		}

		for i := nums[index]; i >= 1; i-- {
			res := recur(index + i)

			if res {
				return res
			}
		}

		memo[index] = false

		return false
	}

	return recur(0)
}
