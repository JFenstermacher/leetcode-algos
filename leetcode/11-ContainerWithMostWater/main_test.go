package main

import "testing"

func TestMaxArea(t *testing.T) {
	inputs := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
		{1, 2, 1},
	}
	outputs := []int{49, 1, 2}

	for i, input := range inputs {
		output := MaxArea(input)

		if outputs[i] != output {
			t.Errorf("Solution %d failed", i)
		}
	}
}
