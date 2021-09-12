package main

import "testing"

func TestReverse(t *testing.T) {
	inputs := []int{123, -321, 1534236469}
	outputs := []int{321, -123, 0}

	for i, input := range inputs {
		output := Reverse(input)

		if output != outputs[i] {
			t.Errorf("Solution %d failed", i)
		}
	}
}
