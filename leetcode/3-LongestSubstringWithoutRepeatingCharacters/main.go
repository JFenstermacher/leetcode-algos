package main

func LengthOfLongestSubstring(s string) int {
	head := 0
	maxLength := 0
	seen := map[rune]int{}

	for i, c := range s {
		last, found := seen[c]

		if found && last >= head {
			head = last + 1
		}

		currLength := i - head + 1

		if currLength > maxLength {
			maxLength = currLength
		}

		seen[c] = i
	}

	return maxLength
}

func main() {
	input := "abba"

	LengthOfLongestSubstring(input)
}
