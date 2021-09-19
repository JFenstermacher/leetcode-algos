package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	l1 := convertToList([]int{2, 4, 3})
	l2 := convertToList([]int{5, 6, 4})

	expected := []int{7, 0, 8}

	res := AddTwoNumbers(l1, l2)

	index := 0
	for res != nil {
		if res.Val != expected[index] {
			t.Error("Solution doesn't work")
		}

		res = res.Next
		index++
	}

	l1 = convertToList([]int{5})
	l2 = convertToList([]int{5})

	expected = []int{0, 1}
	res = AddTwoNumbers(l1, l2)

	index = 0
	for res != nil {
		if res.Val != expected[index] {
			t.Error("Solution doesn't work")
		}

		res = res.Next
		index++
	}
}

func TestCarryOn(t *testing.T) {
	li := convertToList([]int{9, 9, 9})

	res := CarryOn(li, 1)

	expected := []int{0, 0, 0, 1}

	index := 0
	for res != nil {
		if res.Val != expected[index] {
			t.Error("Solution doesn't work")
		}

		res = res.Next
		index++
	}
}
