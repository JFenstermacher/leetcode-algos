package main

import "fmt"

func Reverse(x int) int {
	var sign int32 = 1

	// Max and min values possible before overflow will occur
	max := int32((1<<31 - 1) / 10)
	min := int32((-1 << 31) / 10)

	if x < 0 {
		sign *= -1
	}

	var res int32 = 0

	for x != 0 {
		next := x % 10

		// Check if adding another digit will overflow
		if (sign == 1 && res > max) || (sign == -1 && res < min) {
			return 0
		}

		res *= 10
		res += int32(next)

		x /= 10
	}

	return int(res)
}
