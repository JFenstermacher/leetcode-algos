package main

import "fmt"

func MaxArea(height []int) int {
	maxArea := 0

	for i, j := 0, len(height)-1; i < j; {
		left := height[i]
		right := height[j]
		width := j - i
		height := 0

		if left > right {
			j--
			height = right
		} else {
			i++
			height = left
		}

		area := height * width

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func main() {
	output := MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})

	fmt.Printf("%d", output)
}
