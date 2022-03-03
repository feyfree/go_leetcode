package p0120

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	if n == 1 {
		return triangle[0][0]
	}
	for i := 0; i < n; i++ {
		dp[n-1][i] = triangle[n-1][i]
	}
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if dp[i+1][j] > dp[i+1][j+1] {
				dp[i][j] = dp[i+1][j+1] + triangle[i][j]
			} else {
				dp[i][j] = dp[i+1][j] + triangle[i][j]
			}
		}
	}
	return dp[0][0]
}
