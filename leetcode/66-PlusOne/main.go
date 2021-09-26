package main

import "fmt"

func plusOne(digits []int) []int {
	last := len(digits) - 1
	carry := 0

	digits[last] = (digits[last] + 1) % 10

	if digits[last] == 0 {
		carry = 1
	}

	for i := len(digits) - 2; i >= 0; i-- {
		val := digits[i] + carry

		if val == 10 {
			carry = 1
			digits[i] = 0
		} else {
			carry = 0
			digits[i] = val
		}
	}

	if carry == 1 {
		digits[0]++
		digits = append(digits, 0)
	}

	return digits
}

func main() {
	output := plusOne([]int{9, 9, 9})

	fmt.Println(output)
}
