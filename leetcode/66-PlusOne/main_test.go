package main

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{0},
		{9},
	}
	outputs := [][]int{
		{1, 2, 4},
		{4, 3, 2, 2},
		{1},
		{1, 0},
	}

	for i, input := range inputs {
		output := plusOne(input)

		if !reflect.DeepEqual(output, outputs[i]) {
			t.Errorf("Solution %d failed", i)
		}
	}
}
