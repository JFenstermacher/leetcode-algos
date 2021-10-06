package main

import "testing"

func TestExist(t *testing.T) {
	input := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word := "ABCB"

	if exist(input, word) {
		t.Error("The solution failed")
	}
}
