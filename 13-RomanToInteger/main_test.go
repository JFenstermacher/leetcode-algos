package main

import "testing"

func TestRomanToInt(t *testing.T) {
	inputs := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
	outputs := []int{3, 4, 9, 58, 1994}

	for i, input := range inputs {
		output := romanToInt(input)

		if outputs[i] != output {
			t.Errorf("Solution %d failed", i)
		}
	}
}
