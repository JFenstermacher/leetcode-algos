package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	inputs := []int{-123}
	outputs := []bool{false}

	for i, input := range inputs {
		output := isPalindrome(input)

		if outputs[i] != output {
			t.Errorf("Solution %d failed", i)
		}
	}
}
