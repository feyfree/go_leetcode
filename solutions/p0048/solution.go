package p0048

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		temp := matrix[i]
		matrix[i] = matrix[n-1-i]
		matrix[n-1-i] = temp
	}
	// 对角线 交换
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			temp := matrix[j][i]
			matrix[j][i] = matrix[i][j]
			matrix[i][j] = temp
		}
	}
}
