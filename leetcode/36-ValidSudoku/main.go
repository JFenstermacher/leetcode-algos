package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rows := make([]bool, 81)
	cols := make([]bool, 81)
	squares := make([]bool, 81)

	for i, row := range board {
		for j, cell := range row {
			if cell == '.' {
				continue
			}

			asInt := int(cell-'0') - 1

			rowIndex := i*9 + asInt
			colIndex := j*9 + asInt
			squareIndex := 9*(3*int(i/3)+int(j/3)) + asInt

			if rows[rowIndex] || cols[colIndex] || squares[squareIndex] {
				return false
			}

			rows[rowIndex] = true
			cols[colIndex] = true
			squares[squareIndex] = true
		}
	}

	return true
}

func main() {
	input := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	output := isValidSudoku(input)

	fmt.Print(output)
}
