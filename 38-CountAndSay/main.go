package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	res := []int{1}

	next := func(last []int) []int {
		arr := []int{}

		for i := 0; i < len(last); {
			curr := last[i]
			count := 0

			for ; i < len(last); i, count = i+1, count+1 {
				if curr != last[i] {
					break
				}
			}

			arr = append(arr, count)
			arr = append(arr, curr)
		}

		return arr
	}

	for i := 1; i < n; i++ {
		res = next(res)
	}

	return convertToString(res)
}

func convertToString(arr []int) string {
	var builder strings.Builder

	for _, el := range arr {
		builder.WriteString(strconv.Itoa(el))
	}

	return builder.String()
}

func main() {
	output := countAndSay(1)

	fmt.Println(output)
}
