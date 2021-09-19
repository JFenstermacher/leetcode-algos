package main

import "fmt"

func LongestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	longest := ""

	for i := 0; i < len(s)-1; i++ {
		oddPalindrome := checkPalindrome(s, i, i)
		evenPalindrome := checkPalindrome(s, i, i+1)

		if len(oddPalindrome) > len(longest) {
			longest = oddPalindrome
		}

		if len(evenPalindrome) > len(longest) {
			longest = evenPalindrome
		}
	}

	return longest
}

func checkPalindrome(s string, left int, right int) string {
	curr := ""

	for i, j := left, right; i >= 0 && j < len(s); i, j = i-1, j+1 {
		if s[i] == s[j] {
			curr = s[i : j+1]
		} else {
			break
		}
	}

	return curr
}

func main() {
	output := LongestPalindrome("babad")

	fmt.Println(output)
}
