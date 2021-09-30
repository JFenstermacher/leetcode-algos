package main

import "testing"

func TestUniquePaths(t *testing.T) {
	inputs := [][]int{
		{3, 7},
		{3, 2},
		{7, 3},
		{3, 3},
	}
	outputs := []int{28, 3, 28, 6}

	for i, input := range inputs {
		output := uniquePaths(input[0], input[1])

		if output != outputs[i] {
			t.Errorf("Solution %d failed", i)
		}
	}
}
