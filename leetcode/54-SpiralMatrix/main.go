package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	ordered := []int{}
	rowOffset, colOffset := 1, 1

	count, dir, r, c := len(matrix)*len(matrix[0]), 0, 0, 0
	for count > 0 {
		if dir == 0 {
			for ; c < len(matrix[0])-colOffset; c++ {
				ordered = append(ordered, matrix[r][c])
				count--
			}
		} else if dir == 1 {
			for ; r < len(matrix)-rowOffset; r++ {
				ordered = append(ordered, matrix[r][c])
				count--
			}
		} else if dir == 2 {
			for ; c >= colOffset; c-- {
				ordered = append(ordered, matrix[r][c])
				count--
			}

			colOffset++
		} else {
			for ; r >= rowOffset; r-- {
				ordered = append(ordered, matrix[r][c])
				count--
			}

			r += 1
			c += 1

			rowOffset++
		}

		dir = (dir + 1) % 4
	}

	return ordered
}

func main() {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	output := spiralOrder(input)

	fmt.Println(output)
}
