package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i int, j int) bool {
		iMin, iMax := intervals[i][0], intervals[i][1]
		jMin, jMax := intervals[j][0], intervals[j][1]

		if iMin < jMin {
			return true
		} else if iMin == 0 {
			return iMax <= jMax
		} else {
			return false
		}
	})

	res := [][]int{}
	curr := []int{intervals[0][0], intervals[0][1]}

	for i := 1; i < len(intervals); i++ {
		elMin, elMax := intervals[i][0], intervals[i][1]

		if elMin <= curr[1] {
			if curr[1] < elMax {
				curr[1] = elMax
			}
		} else {
			res = append(res, append([]int{}, curr...))
			curr[0] = elMin
			curr[1] = elMax
		}
	}

	res = append(res, append([]int{}, curr...))

	return res
}

func main() {
	arr := [][]int{
		{1, 4},
		{4, 5},
	}

	output := merge(arr)

	fmt.Println(output)
}
