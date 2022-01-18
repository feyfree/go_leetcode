package p0073

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	col := false
	row := false
	for i := 0; i < m; i++ {
		col = col || matrix[i][0] == 0
	}
	for i := 0; i < n; i++ {
		row = row || matrix[0][i] == 0
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if col {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

}
