package main

import "testing"

func compareArrs(a1 []string, a2 []string) bool {
	if len(a1) != len(a2) {
		return false
	}

	mapping := map[string]interface{}{}

	for _, val := range a1 {
		mapping[val] = nil
	}

	for _, val := range a2 {
		_, found := mapping[val]

		if !found {
			return false
		}
	}

	return true
}

func TestGenerateParentheses(t *testing.T) {
	inputs := []int{3, 1}
	outputs := [][]string{
		{"((()))", "(()())", "(())()", "()(())", "()()()"},
		{"()"},
	}

	for i, input := range inputs {
		output := generateParentheses(input)

		if !compareArrs(output, outputs[i]) {
			t.Errorf("Solution %d failed", i)
		}
	}
}
