package main

import "testing"

func TestStrStr(t *testing.T) {
	inputs := [][]string{
		{"hello", "ll"},
		{"aaaaa", "bba"},
		{"", ""},
		{"", "a"},
		{"a", "a"},
		{"aaa", "aaaa"},
		{"mississippi", "issip"},
	}

	outputs := []int{2, -1, 0, -1, 0, -1, 4}

	for i, input := range inputs {
		output := strStr(input[0], input[1])

		if output != outputs[i] {
			t.Errorf("Solution %d failed", i)
		}
	}
}
