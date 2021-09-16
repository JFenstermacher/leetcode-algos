package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	inputs := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	}
	outputs := []int{2, 5}

	for i, input := range inputs {
		output := removeDuplicates(input)

		if outputs[i] != output {
			t.Errorf("Solution %d failed", i)
		}
	}
}
