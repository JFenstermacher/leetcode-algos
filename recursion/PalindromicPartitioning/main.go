package main

import "fmt"

func isPalindrome(s string) bool {
	for low, high := 0, len(s)-1; low < high; low, high = low+1, high-1 {
		if s[low] != s[high] {
			return false
		}
	}

	return true
}

func palindromePartitioning(s string) [][]string {
	results := [][]string{}

	var recur func(index int, combo []string)

	recur = func(start int, combo []string) {
		if start == len(s) {
			copied := make([]string, len(combo))
			copy(copied, combo)

			results = append(results, copied)
		}

		for i := start; i < len(s); i++ {
			curr := s[start : i+1]
			if isPalindrome(curr) {
				combo = append(combo, curr)
				recur(i+1, combo)
				combo = combo[:len(combo)-1]
			}
		}
	}

	recur(0, []string{})

	return results
}

func main() {
	input := "geeks"

	output := palindromePartitioning(input)

	fmt.Println(output)
}
