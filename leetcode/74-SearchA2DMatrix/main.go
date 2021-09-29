package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	r, c := 0, 0

	for i, row := range matrix {
		first, last := row[0], row[len(row)-1]

		if first <= target && target <= last {
			r = i
			c = binarySearch(row, target)
		}
	}

	return matrix[r][c] == target
}

func binarySearch(row []int, target int) int {
	low, high := 0, len(row)-1

	mid := 0
	for low <= high {
		mid = (low + high) / 2

		curr := row[mid]
		if curr == target {
			break
		} else if curr < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return mid
}

func main() {
	arr := [][]int{
		{1, 3},
	}

	output := searchMatrix(arr, 3)

	fmt.Println(output)
}
