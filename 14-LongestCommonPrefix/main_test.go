package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	inputs := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{"ab", "a"},
	}
	outputs := []string{"fl", "", "a"}

	for i, input := range inputs {
		output := longestCommonPrefix(input)

		if outputs[i] != output {
			t.Errorf("Solution %d failed", i)
		}
	}
}
