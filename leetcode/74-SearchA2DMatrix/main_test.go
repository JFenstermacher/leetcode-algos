package main

import "testing"

func TestSearchMatrix(t *testing.T) {
	arr := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}

	targets := []int{3, 13}

	outputs := []bool{true, false}

	for i, target := range targets {
		output := searchMatrix(arr, target)

		if output != outputs[i] {
			t.Errorf("Solution %d failed", i)
		}
	}
}
