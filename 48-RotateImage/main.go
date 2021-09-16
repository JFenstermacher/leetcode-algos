package main

import "fmt"

func rotate(matrix [][]int) [][]int {
	matrix = foldY(matrix)

	return foldXY(matrix)
}

func foldY(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < int(len(matrix)/2); j++ {
			tmp := matrix[i][j]

			matrix[i][j] = matrix[i][len(matrix)-1-j]

			matrix[i][len(matrix)-1-j] = tmp
		}
	}

	return matrix
}

func foldXY(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)-1-i; j++ {
			r := (len(matrix) - 1 - j)
			c := (len(matrix) - 1 - i)

			tmp := matrix[i][j]
			matrix[i][j] = matrix[r][c]
			matrix[r][c] = tmp
		}
	}

	return matrix
}

func main() {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	output := foldXY(input)

	fmt.Println(output)
}
