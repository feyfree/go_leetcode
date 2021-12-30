package p0044

// s 是给定字符串
// p 是pattern
func isMatch(s string, p string) bool {
	l1 := len(s)
	l2 := len(p)
	dp := make([][]bool, l1+1)
	for i := range dp {
		dp[i] = make([]bool, l2+1)
	}
	dp[0][0] = true
	// 初始化
	for i := 1; i <= l2; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			sc := s[i-1]
			pc := p[j-1]
			if pc == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else if pc == '?' || sc == pc {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[l1][l2]
}
