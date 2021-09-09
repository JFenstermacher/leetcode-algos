package main

import "strconv"

func isPalindrome(x int) bool {
	asString := strconv.Itoa(x)

	for i, c := range asString {
		back := len(asString) - i - 1

		if byte(c) != asString[back] {
			return false
		}
	}

	return true
}
