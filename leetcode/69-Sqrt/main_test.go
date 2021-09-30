package main

import "testing"

func TestMySqrt(t *testing.T) {
	inputs := []int{4, 8}
	outputs := []int{2, 2}

	for i, input := range inputs {
		output := mySqrt(input)

		if output != outputs[i] {
			t.Errorf("Solution %d failed", i)
		}
	}
}
