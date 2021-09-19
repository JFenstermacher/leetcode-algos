package main

func canJumpSlow(nums []int) bool {
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

func canJump(nums []int) bool {
	max := 0

	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		} else {
			currMax := i + nums[i]

			if currMax > max {
				max = currMax
			}
		}
	}

	return true
}
