package main

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	inputs := [][]int{
		{2, 3, 6, 7, 7},
		{2, 3, 5, 8},
		{2, 1},
		{1, 1},
	}
	outputs := [][][]int{
		{{2, 2, 3}, {7}},
		{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		{},
		{{1}},
	}

	for i, input := range inputs {
		candidates, target := input[:len(input)-1], input[len(input)-1]

		output := combinationSum(candidates, target)

		if !reflect.DeepEqual(output, outputs[i]) {
			t.Errorf("Solution %d failed", i)
		}
	}
}
