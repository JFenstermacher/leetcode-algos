package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	results := [][]int{}

	var recur func(remain int, index int, combo []int)

	recur = func(remain int, index int, combo []int) {
		if remain == 0 {
			copied := make([]int, len(combo))
			copy(copied, combo)

			results = append(results, copied)
			return
		} else if remain < 0 {
			return
		}

		for i, num := range candidates[index:] {
			combo = append(combo, num)
			recur(remain-num, index+i, combo)
			combo = combo[:len(combo)-1]
		}

		return
	}

	recur(target, 0, []int{})

	return results
}

func main() {
	candidates := []int{2, 7, 6, 3, 5, 1}
	target := 9

	output := combinationSum(candidates, target)

	fmt.Println(output)
}
