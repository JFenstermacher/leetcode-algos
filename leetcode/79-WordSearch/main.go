package main

import "fmt"

func exist(board [][]byte, word string) bool {
	cols := len(board[0])
	var recur func(r int, c int, index int, used int64) bool

	calc := func(r int, c int) int {
		return r*cols + c
	}

	recur = func(r int, c int, index int, used int64) bool {
		if index >= len(word) {
			return true
		}

		if r < 0 || r == len(board) || c < 0 || c == len(board[0]) {
			return false
		}

		var key int64 = 1 << calc(r, c)

		if key&used > 0 {
			return false
		}

		if word[index] == board[r][c] {
			used += key
			return recur(r-1, c, index+1, used) || recur(r+1, c, index+1, used) || recur(r, c-1, index+1, used) || recur(r, c+1, index+1, used)
		} else {
			return false
		}
	}

	for i := range board {
		for j := range board[i] {
			used := make([][]bool, len(board))
			for i := 0; i < len(used); i++ {
				used[i] = make([]bool, len(board[0]))
			}

			if recur(i, j, 0, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	input := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCESEEEFS"

	output := exist(input, word)

	fmt.Println(output)
}
