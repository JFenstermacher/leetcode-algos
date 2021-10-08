package main

import "testing"

func TestPartition(t *testing.T) {
	li := convertToList([]int{1, 4, 3, 2, 5, 2})
	output := partition(li, 3)

	expected := []int{1, 2, 2, 4, 3, 5}

	for i := 0; output != nil; i++ {
		if expected[i] != output.Val {
			t.Errorf("Solution fails")

			break
		}

		output = output.Next
	}
}
