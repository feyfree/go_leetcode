package p0005

// 动态规划问题
func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	left, length := 0, 1
	// 相当于是end
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		// 相当于是start
		for j := 0; j < i; j++ {
			if i-j == 1 {
				dp[j][i] = s[j] == s[i]
			} else {
				dp[j][i] = s[j] == s[i] && dp[j+1][i-1]
			}
			if dp[j][i] && i-j+1 > length {
				length = i - j + 1
				left = j
			}
		}
	}
	return s[left : left+length]
}
