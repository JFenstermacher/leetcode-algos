package main

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	inputs := []string{
		"Hello world",
		"   fly me   to   the moon    ",
		"luffy is still joyboy",
	}
	outputs := []int{5, 4, 6}

	for i, input := range inputs {
		output := lengthOfLastWord(input)

		if outputs[i] != output {
			t.Errorf("Solution %d failed", i)
		}
	}
}
