package main

import (
	"reflect"
	"testing"
)

func TestMultiplyOne(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 4, 5, 0},
		{1, 2, 3, 4, 5, 1},
		{3, 2, 1, 4, 0},
	}
	outputs := [][]int{
		{5, 0, 6, 1, 2},
		{0, 5, 0, 6, 1, 2},
		{2, 9, 4},
	}

	for i, input := range inputs {
		arr, multi, padding := input[:len(input)-2], input[len(input)-2], input[len(input)-1]

		output := multiplyOne(arr, multi, padding)

		if !reflect.DeepEqual(output, outputs[i]) {
			t.Errorf("MultiplyOne fails on %d input", i)
		}
	}
}

func TestAdd(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{1, 2, 7},
	}

	expected := []int{2, 4, 0, 1}

	output := add(input[0], input[1])

	if !reflect.DeepEqual(expected, output) {
		t.Error("Add test failed")
	}
}

func TestConvertArrToString(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}

	expected := "7654321"

	output := convertArrToString(input)

	if expected != output {
		t.Error("Convert To String failed")
	}
}

func TestMultiply(t *testing.T) {
	inputs := [][]string{
		{"2", "3"},
		{"123", "456"},
		{"9133", "0"},
	}

	outputs := []string{"6", "56088", "0"}

	for i, input := range inputs {
		output := multiply(input[0], input[1])

		if output != outputs[i] {
			t.Errorf("Multiply failed for test case %d", i)
		}
	}
}
