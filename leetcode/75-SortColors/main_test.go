package main

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	inputs := [][]int{
		{2, 0, 2, 1, 1, 0},
		{2, 0, 1},
		{0},
		{1},
	}
	outputs := [][]int{
		{0, 0, 1, 1, 2, 2},
		{0, 1, 2},
		{0},
		{1},
	}

	for i, input := range inputs {
		sortColors(input)

		if !reflect.DeepEqual(input, outputs[i]) {
			t.Errorf("Solution %d failed", i)
		}
	}
}
