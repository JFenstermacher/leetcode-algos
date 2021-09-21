package main

func maxSubArray(nums []int) int {
	sum, min, res := 0, 0, nums[0]

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if sum-min > res {
			res = sum - min
		}
		if sum < min {
			min = sum
		}
	}

	return res
}
