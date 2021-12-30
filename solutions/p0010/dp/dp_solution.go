package dp

func isMatch(s string, p string) bool {
	// 动态规划 二维数组
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[m][n] = true
	for i := m; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// i, j 的字面匹配
			firstMatch := i < m && (s[i] == p[j] || p[j] == '.')
			// 更新dp[i][j]
			if j < n-1 && p[j+1] == '*' {
				dp[i][j] = (firstMatch && (dp[i+1][j] || dp[i][j+2])) || (!firstMatch && dp[i][j+2])
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}
