package main

import "fmt"

func uniquePaths(m int, n int) int {
	arr := make([][]int, m)

	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
		arr[i][0] = 1
	}

	for i := 0; i < n; i++ {
		arr[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			arr[i][j] = arr[i-1][j] + arr[i][j-1]
		}
	}

	return arr[m-1][n-1]
}

func main() {
	output := uniquePaths(3, 3)

	fmt.Println(output)
}
