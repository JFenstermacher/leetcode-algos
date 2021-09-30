package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
			break
		} else {
			obstacleGrid[i][0] = 1
		}
	}

	for i := 1; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			obstacleGrid[0][i] = 0
			break
		} else {
			obstacleGrid[0][i] = 1
		}
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}

			obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
		}
	}

	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func main() {
	output := uniquePathsWithObstacles([][]int{
		{0, 1},
		{0, 0},
	})

	fmt.Println(output)
}
