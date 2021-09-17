package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	mapping := map[string]interface{}{}

	for _, str := range strs {
		counts := make([]int, 26)

		for _, c := range str {
			counts[c-97]++
		}

		key := fmt.Sprint(counts)

		arr, found := mapping[key]

		if !found {
			mapping[key] = []string{str}
		} else {
			arr = append(arr.([]string), str)
			mapping[key] = arr
		}
	}

	res := [][]string{}
	for _, val := range mapping {
		res = append(res, val.([]string))
	}

	return res
}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	output := groupAnagrams(input)

	fmt.Print(output)
}
