package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	path = strings.ReplaceAll(path, "//", "/")
	parts := strings.Split(path, "/")
	stack := []string{}

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if part != "." {
			stack = append(stack, part)
		}
	}

	return fmt.Sprintf("/%s", strings.Join(stack, "/"))
}

func main() {
	output := simplifyPath("/home/")

	fmt.Println(output)
}
