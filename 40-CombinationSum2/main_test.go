package main

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	inputs := [][]int{
		{10, 1, 2, 7, 6, 1, 5, 8},
		{2, 5, 2, 1, 2, 5},
	}
	outputs := [][][]int{
		{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		{{1, 2, 2}, {5}},
	}

	for i, input := range inputs {
		candidates, target := input[:len(input)-1], input[len(input)-1]

		output := combinationSum2(candidates, target)

		if !reflect.DeepEqual(output, outputs[i]) {
			t.Errorf("Solution %d failed", i)
		}
	}
}
