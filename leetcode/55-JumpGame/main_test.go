package main

import "testing"

func TestCanJump(t *testing.T) {
	inputs := [][]int{
		{2, 3, 1, 1, 4},
		{3, 2, 1, 0, 4},
	}

	outputs := []bool{true, false}

	for i, input := range inputs {
		output := canJump(input)

		if output != outputs[i] {
			t.Errorf("Can Jump failed for test case %d", i)
		}
	}
}
