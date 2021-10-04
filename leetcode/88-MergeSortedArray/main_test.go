package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	inputs := [][][]int{
		{
			{1, 1},
			{},
		},
		{
			{1, 2, 3, 0, 0, 0, 3},
			{2, 5, 6},
		},
	}

	outputs := [][]int{
		{1, 2, 2, 3, 5, 6},
		{1},
	}

	for i, input := range inputs {
		li1 := input[0]
		last := len(li1) - 1

		mergeRealSolution(li1[:last], li1[last], input[1], len(input[1]))

		if reflect.DeepEqual(li1[:last], outputs[i]) {
			t.Errorf("Solution %d failed", i)
		}
	}
}
