package main

import "testing"

func TestSearchInsert(t *testing.T) {
	inputs := [][]int{
		{1, 3, 5, 6, 5},
		{1, 3, 5, 6, 2},
		{1, 3, 5, 6, 7},
		{1, 3, 5, 6, 0},
		{1, 0},
	}
	outputs := []int{2, 1, 4, 0, 0}

	for i, input := range inputs {
		last := len(input) - 1
		output := searchInsert(input[:last], input[last])

		if outputs[i] != output {
			t.Errorf("Solution %d failed", i)
		}
	}
}
