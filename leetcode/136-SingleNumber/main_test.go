package main

import "testing"

func TestSingleNumber(t *testing.T) {
	inputs := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
		{1},
	}
	outputs := []int{1, 4, 1}

	for i, input := range inputs {
		output := singleNumber(input)

		if output != outputs[i] {
			t.Errorf("Solution %d failed", i)
		}
	}
}
