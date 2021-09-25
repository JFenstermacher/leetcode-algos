package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	inputs := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {2, 3}},
	}
	outputs := [][][]int{
		{{1, 6}, {8, 10}, {15, 18}},
		{{1, 5}},
		{{1, 4}},
	}

	for i, input := range inputs {
		output := merge(input)

		if !reflect.DeepEqual(output, outputs[i]) {
			t.Errorf("Solution %d failed", i)
		}
	}
}
