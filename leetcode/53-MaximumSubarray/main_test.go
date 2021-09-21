package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	inputs := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1},
		{5, 4, -1, 7, 8},
	}
	outputs := []int{6, 1, 23}

	for i, input := range inputs {
		output := maxSubArray(input)

		if outputs[i] != output {
			t.Errorf("Solution %d failed", i)
		}
	}
}
