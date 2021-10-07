package main

import "fmt"

func generate(numRows int) [][]int {
	res := [][]int{{1}}

	if numRows == 1 {
		return res
	}

	res = append(res, []int{1, 1})

	if numRows == 2 {
		return res
	}

	for i := 2; i < numRows; i++ {
		row := []int{1}

		last := res[len(res)-1]
		for j := 0; j < len(last)-1; j++ {
			row = append(row, last[j]+last[j+1])
		}

		row = append(row, 1)
		res = append(res, row)
	}

	return res
}

func main() {
	output := generate(5)

	fmt.Println(output)
}
