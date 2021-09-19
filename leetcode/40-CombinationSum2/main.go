package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	combos := [][]int{}

	var recur func(remain int, index int, combo []int)

	recur = func(remain int, index int, combo []int) {
		if remain == 0 {
			combos = append(combos, combo)
			return
		}

		if remain < 0 {
			return
		}

		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}

			newCombo := append([]int{}, combo...)
			newCombo = append(newCombo, candidates[i])
			recur(remain-candidates[i], i+1, newCombo)
		}
	}

	recur(target, 0, []int{})

	return combos
}

func combinationSum2b(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	combos := [][]int{}

	var recur func(remain int, index int, combo []int)

	recur = func(remain int, index int, combo []int) {
		fmt.Println(remain, index, combo)
		if remain == 0 {
			combos = append(combos, append([]int{}, combo...))
			return
		}

		if index == len(candidates) || remain < 0 {
			return
		}

		recur(remain, index+1, combo)
		val := candidates[index]
		if index > 0 && candidates[index-1] == val && remain == target {
			return
		}

		combo = append(combo, val)
		recur(remain-val, index+1, combo)
		combo = combo[:len(combo)-1]
	}

	recur(target, 0, []int{})

	return combos
}

func main() {
	input := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	output := combinationSum2b(input, target)

	fmt.Println(output)
}
