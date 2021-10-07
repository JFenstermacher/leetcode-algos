package main

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	inputs := []int{5, 1}
	outputs := [][][]int{
		{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		{{1}},
	}

	for i, input := range inputs {
		output := generate(input)

		if !reflect.DeepEqual(output, outputs[i]) {
			t.Errorf("Solution %d failed", i)
		}
	}
}
