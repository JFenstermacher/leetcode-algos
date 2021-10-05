package main

import "testing"

func TestMinPathSum(t *testing.T) {
	inputs := [][][]int{
		{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		},
		{
			{1, 2, 3},
			{4, 5, 6},
		},
	}
	outputs := []int{7, 12}

	for i, input := range inputs {
		output := minPathSum(input)

		if output != outputs[i] {
			t.Errorf("Solution %d failed", i)
		}
	}
}
