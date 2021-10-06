package main

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3},
		{1, 2},
	}
	outputs := [][][]int{
		{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		{{}, {1}, {2}, {1, 2}},
	}

	for i, input := range inputs {
		output := subsets(input)

		if !reflect.DeepEqual(output, outputs[i]) {
			t.Errorf("Solution %d failed", i)
		}
	}
}
