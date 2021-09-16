package main

import "testing"

func TestRemoveElement(t *testing.T) {
	inputs := [][]int{
		{3, 2, 2, 3, 3},
		{0, 1, 2, 2, 3, 0, 4, 2, 2},
	}
	outputs := []int{2, 5}

	for i, input := range inputs {
		length := len(input)
		output := removeElement(input[:length-1], input[length-1])

		if outputs[i] != output {
			t.Errorf("Solution %d failed", i)
		}
	}
}
