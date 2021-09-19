package main

import "testing"

func TestIsValid(t *testing.T) {
	inputs := [][][]int{
		{{1, 2, 4}, {1, 3, 4}},
		{{}, {}},
		{{}, {0}},
	}
	outputs := []*ListNode{
		convertToList([]int{1, 1, 2, 3, 4, 4}),
		convertToList([]int{}),
		convertToList([]int{0}),
	}

	for i, input := range inputs {
		l1 := convertToList(input[0])
		l2 := convertToList(input[1])

		output := mergeTwoLists(l1, l2)

		if !compareLists(output, outputs[i]) {
			t.Errorf("Solution %d failed", i)
		}
	}
}
