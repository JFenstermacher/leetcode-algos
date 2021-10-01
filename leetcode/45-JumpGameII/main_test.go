package main

import "testing"

func TestJump(t *testing.T) {
	inputs := [][]int{
		{2, 3, 1, 1, 4},
		{2, 3, 0, 1, 4},
	}
	outputs := []int{2, 2}

	for i, input := range inputs {
		output := jump(input)

		if output != outputs[i] {
			t.Errorf("Solution %d failed", i)
		}
	}
}
