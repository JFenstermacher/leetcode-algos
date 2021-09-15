package main

import "testing"

func TestIsValid(t *testing.T) {
	inputs := []string{"()", "()[]{}", "(]", "([)]", "{[]}", "[", "]"}
	outputs := []bool{true, true, false, false, true, false, false}

	for i, input := range inputs {
		output := isValid(input)

		if outputs[i] != output {
			t.Errorf("Solution %d failed", i)
		}
	}
}
