package main

import "testing"

func compareMatrixes(m1 [][]int, m2 [][]int) bool {
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m1); j++ {
			if m1[i][j] != m2[i][j] {
				return false
			}
		}
	}

	return true
}

func TestFoldY(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := [][]int{
		{3, 2, 1},
		{6, 5, 4},
		{9, 8, 7},
	}

	output := foldY(input)

	if !compareMatrixes(output, expected) {
		t.Error("Flip Y failed")
	}
}

func TestFoldXY(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := [][]int{
		{9, 6, 3},
		{8, 5, 2},
		{7, 4, 1},
	}

	output := foldXY(input)

	if !compareMatrixes(output, expected) {
		t.Error("Flip XY failed")
	}
}

func TestRotate(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	output := rotate(input)

	if !compareMatrixes(output, expected) {
		t.Error("Flip XY failed")
	}
}
