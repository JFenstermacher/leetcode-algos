package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		count := 0
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] == needle[j] {
				count++
			}
		}

		if count == len(needle) {
			return i
		}
	}

	return -1
}

func main() {
	output := strStr("mississippi", "issip")

	fmt.Println(output)
}
