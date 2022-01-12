package p0063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	return dfs(0, 0, m, n, &dp, &obstacleGrid)
}

func dfs(x int, y int, m int, n int, dp *[][]int, obstacleGrid *[][]int) int {
	if x > m-1 || y > n-1 {
		return 0
	}
	if (*obstacleGrid)[x][y] == 1 {
		return 0
	}
	if x == m-1 && y == n-1 {
		return 1
	}
	if (*dp)[x][y] == 0 {
		(*dp)[x][y] = dfs(x+1, y, m, n, dp, obstacleGrid) + dfs(x, y+1, m, n, dp, obstacleGrid)
	}
	return (*dp)[x][y]
}
