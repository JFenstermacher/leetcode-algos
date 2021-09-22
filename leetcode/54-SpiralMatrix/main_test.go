package main

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}

	output := spiralOrder(input)

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Spiral Order solution failed")
	}
}
