package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	inputs := []string{"abcabcbb", " ", "aab", "dvdf", "abba"}
	outputs := []int{3, 1, 2, 3, 2}

	for i, input := range inputs {
		output := LengthOfLongestSubstring(input)

		if output != outputs[i] {
			t.Errorf("Solution %d is incorrect", i)
		}
	}
}
