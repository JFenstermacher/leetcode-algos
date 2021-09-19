package main

import (
	"fmt"
	"strconv"
	"strings"
	// "strconv"
)

func multiply(num1 string, num2 string) string {
	numInts := []int{}
	res := []int{}

	for i := len(num1) - 1; i >= 0; i-- {
		numInts = append(numInts, int(num1[i]-'0'))
	}

	for i, c := range num2 {
		curr := int(c - '0')

		multiplied := multiplyOne(numInts, curr, len(num2)-1-i)
		res = add(res, multiplied)
	}

	asString := convertArrToString(res)

	if asString[0] == '0' {
		return "0"
	}

	return asString
}

func convertArrToString(nums []int) string {
	var builder strings.Builder

	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		builder.WriteString(strconv.Itoa(num))
	}

	return builder.String()
}

func multiplyOne(nums []int, multi int, padding int) []int {
	res := make([]int, padding)
	carry := 0

	for _, num := range nums {
		val := num*multi + carry

		if val >= 10 {
			carry = val / 10
			val -= (10 * carry)
		} else {
			carry = 0
		}

		res = append(res, val)
	}

	if carry > 0 {
		res = append(res, carry)
	}

	return res
}

// Adds two arrays
func add(nums1 []int, nums2 []int) []int {
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	index := 0
	res := []int{}
	carry := 0

	for ; index < len(nums1); index++ {
		num1, num2 := nums1[index], nums2[index]

		sum := num1 + num2 + carry

		if sum >= 10 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}

		res = append(res, sum)
	}

	for ; index < len(nums2); index++ {
		num := nums2[index] + carry

		if num >= 10 {
			carry = 1
			num -= 10
		} else {
			carry = 0
		}

		res = append(res, num)
	}

	if carry == 1 {
		res = append(res, 1)
	}

	return res
}

func main() {
	output := multiply("123", "456")

	first := 123 * 4
	second := 123 * 50
	third := 123 * 600

	total := first + second + third

	fmt.Println(first, second, third, total)

	fmt.Println(output)
}
