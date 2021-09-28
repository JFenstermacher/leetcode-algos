package main

import "testing"

func TestSimplifyPath(t *testing.T) {
	inputs := []string{"/home/", "/../", "/home//foo", "/a/./b/../../c/"}
	outputs := []string{"/home", "/", "/home/foo", "/c"}

	for i, input := range inputs {
		output := simplifyPath(input)

		if output != outputs[i] {
			t.Errorf("Solution %d failed", i)
		}
	}
}
