package p0115

// 计算 s 在 t 中的出现次数
func numDistinct(s string, t string) int {
	lt := len(t)
	ls := len(s)
	dp := make([][]int, lt+1)
	for i := 0; i < lt+1; i++ {
		dp[i] = make([]int, ls+1)
	}
	// 初始化
	for i := 0; i < ls+1; i++ {
		dp[0][i] = 1
	}
	for i := 1; i <= lt; i++ {
		for j := 1; j <= ls; j++ {
			// s 是 target
			// t 是 搜索范围
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[lt][ls]
}
