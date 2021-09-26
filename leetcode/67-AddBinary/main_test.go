package main

import "testing"

func TestAddBinary(t *testing.T) {
	inputs := [][]string{
		{"11", "1"},
		{"1010", "1011"},
	}
	outputs := []string{
		"100",
		"10101",
	}

	for i, input := range inputs {
		output := addBinary(input[0], input[1])

		if output != outputs[i] {
			t.Errorf("Solution failed on test case #%d", i)
		}
	}
}
