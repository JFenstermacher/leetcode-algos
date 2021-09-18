package main

import "testing"

func compareToArray(node *ListNode, arr []int) bool {
	index := 0

	for node != nil {
		if node.Val != arr[index] || index >= len(arr) {
			return false
		}

		index++
		node = node.Next
	}

	return index == len(arr)
}

func TestSwapPairs(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 4},
		{},
		{1},
	}

	outputs := [][]int{
		{2, 1, 4, 3},
		{},
		{1},
	}

	for i, input := range inputs {
		output := swapPairs(convertToList(input))

		if !compareToArray(output, outputs[i]) {
			t.Errorf("Solution failed for %d", i)
		}
	}
}

func TestRecursiveSwapPairs(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 4},
		{},
		{1},
	}

	outputs := [][]int{
		{2, 1, 4, 3},
		{},
		{1},
	}

	for i, input := range inputs {
		output := recursiveSwapPairs(convertToList(input))

		if !compareToArray(output, outputs[i]) {
			t.Errorf("Solution failed for %d", i)
		}
	}
}
