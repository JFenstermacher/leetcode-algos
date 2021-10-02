package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	converted := convert(s)
	if len(converted) == 1 {
		return true
	}

	mid := int(len(converted) / 2)

	if len(converted)%2 == 0 {
		return checkPalindrome(converted, mid-1, mid)
	} else {
		return checkPalindrome(converted, mid, mid)
	}
}

func convert(s string) string {
	lower := strings.ToLower(s)

	regex, _ := regexp.Compile("[^a-z0-9]")

	return regex.ReplaceAllLiteralString(lower, "")
}

func checkPalindrome(s string, left int, right int) bool {
	for i, j := left, right; i >= 0 && j < len(s); i, j = i-1, j+1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func main() {
	output := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(output)
}
