package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var builder strings.Builder

	for i := 0; i < len(strs[0]); i++ {
		curr := rune(strs[0][i])
		match := true

		for j := 0; j < len(strs); j++ {
			if i == len(strs[j]) || rune(strs[j][i]) != curr {
				match = false
				break
			}
		}

		if !match {
			break
		}

		builder.WriteRune(curr)
	}

	return builder.String()
}

func main() {
	output := longestCommonPrefix([]string{"ab", "a"})

	fmt.Printf("Output: %s", output)
}
