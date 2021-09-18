package main

import "testing"

func TestCountAndSay(t *testing.T) {
	inputs := []int{1, 4, 5, 6}
	outputs := []string{"1", "1211", "111221", "312211"}

	for i, input := range inputs {
		output := countAndSay(input)

		if outputs[i] != output {
			t.Errorf("Solution %d failed", i)
		}
	}
}
