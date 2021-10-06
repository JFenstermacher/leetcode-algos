package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	inputs := [][]int{
		{1, 1, 1, 2, 2, 3},
		{0, 0, 1, 1, 1, 1, 2, 3, 3},
	}
	outputs := []int{5, 7}

	for i, input := range inputs {
		output := removeDuplicates(input)

		if !reflect.DeepEqual(output, outputs[i]) {
			t.Errorf("Solution %d failed", i)
		}
	}
}
