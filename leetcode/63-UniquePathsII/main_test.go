package main

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	inputs := [][][]int{
		{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		{
			{0, 1},
			{0, 0},
		},
	}
	outputs := []int{2, 1}

	for i, input := range inputs {
		output := uniquePathsWithObstacles(input)

		if output != outputs[i] {
			t.Errorf("Solution %d failed", i)
		}
	}
}
