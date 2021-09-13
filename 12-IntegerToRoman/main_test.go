package main

import "testing"

func TestIntToRoman(t *testing.T) {
	inputs := []int{3, 4, 9}
	outputs := []string{"III", "IV", "IX"}

	for i, input := range inputs {
		output := IntToRoman(input)

		if outputs[i] != output {
			t.Errorf("Solution %d failed", i)
		}
	}
}
