package main

import "fmt"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)

	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, n)
	}

	// Write first row
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	// Write first column
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			curr := grid[i][j]
			dp[i][j] = min(curr+dp[i-1][j], curr+dp[i][j-1])
		}
	}

	return dp[m-1][n-1]
}

func min(lhs int, rhs int) int {
	if lhs < rhs {
		return lhs
	}

	return rhs
}

func main() {
	// grid := [][]int{
	// 	{1, 3, 1},
	// 	{1, 5, 1},
	// 	{4, 2, 1},
	// }
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	output := minPathSum(grid)

	fmt.Println(output)
}
