package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if length > 0 {
				break
			}

			continue
		}

		length++
	}

	return length
}

func main() {
	output := lengthOfLastWord("luffy is still joyboy")

	fmt.Println(output)
}
