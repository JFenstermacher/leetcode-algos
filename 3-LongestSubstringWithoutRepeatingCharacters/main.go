package main

import "fmt"

func LengthOfLongestSubstring(s string) int {
	currLength := 0
	maxLength := 0
	seen := make([]bool, 256)

	for i := range s {
		if seen[s[i]] {
			seen = make([]bool, 256)

			currLength = 1
		} else {
			currLength++

			if currLength > maxLength {
				maxLength = currLength
			}
		}

		seen[s[i]] = true
	}

	return maxLength
}

func main() {
	input := "abcabcbb"

	output := LengthOfLongestSubstring(input)

	fmt.Printf("%d", output)
}
