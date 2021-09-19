package main

import "fmt"

func allPossibleStrings(set []rune, k int) []string {
	result := []string{}

	var recur func(arr []rune)

	recur = func(arr []rune) {
		if len(arr) == k {
			result = append(result, string(arr))
			return
		}

		for _, c := range set {
			arr = append(arr, c)
			recur(arr)
			arr = arr[:len(arr)-1]
		}
	}

	recur([]rune{})

	return result
}

func main() {
	output := allPossibleStrings([]rune{'a', 'b', 'c', 'd'}, 4)

	for _, o := range output {
		fmt.Println(o)
	}
}
