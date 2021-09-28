package main

import "testing"

func TestClimbStairs(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7}
	outputs := []int{1, 2, 3, 5, 8, 13, 21}

	for i, input := range inputs {
		output := climbStairs(input)

		if output != outputs[i] {
			t.Errorf("Solution %d failed", i)
		}
	}
}
