package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	inputs := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
	}

	outputs := []bool{true, false}

	for i, input := range inputs {
		output := isPalindrome(input)

		if output != outputs[i] {
			t.Errorf("Solution %d failed", i)
		}
	}
}
