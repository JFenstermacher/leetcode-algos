package main

import (
	"fmt"
	"testing"
)

func compareArrs(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	set := map[string]interface{}{}
	for _, sum := range a {
		s := fmt.Sprintf("%d %d %d", sum[0], sum[1], sum[2])

		set[s] = nil
	}

	for _, sum := range b {
		s := fmt.Sprintf("%d %d %d", sum[0], sum[1], sum[2])

		_, found := set[s]

		if !found {
			return false
		}
	}

	return true
}

func TestThreeSum(t *testing.T) {
	inputs := [][]int{
		{-1, 0, 1, 2, -1, 4},
		{},
		{0},
	}
	outputs := [][][]int{
		{{-1, -1, 2}, {-1, 0, 1}},
		{},
		{},
	}

	for i, input := range inputs {
		output := threeSum(input)

		if !compareArrs(output, outputs[i]) {
			t.Errorf("Solution %d failed", i)
		}
	}
}
