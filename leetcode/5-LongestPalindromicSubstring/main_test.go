package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	inputs := []string{"babad", "a", "", "aacabdkacaa"}
	outputs := []string{"bab", "a", "", "aca"}

	for i, input := range inputs {
		palindrome := LongestPalindrome(input)

		if palindrome != outputs[i] {
			t.Errorf("Solution %d doesn't work", i)
		}
	}
}
