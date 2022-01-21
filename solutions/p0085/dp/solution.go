package dp

import "math"

func maximalRectangle(matrix [][]byte) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	// dp[i][j]  是 max length of all 1 sequence ends with col j, at the i-th row
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				if j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i][j-1] + 1
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	var result int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			width := math.MaxInt32
			for k := i; k < rows; k++ {
				if width > dp[k][j] {
					width = dp[k][j]
				}
				area := width * (k - i + 1)
				if area > result {
					result = area
				}
			}
		}
	}
	return result
}
