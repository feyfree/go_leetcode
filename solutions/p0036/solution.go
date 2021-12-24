package p0036

func isValidSudoku(board [][]byte) bool {
	row := [9][9]bool{}
	col := [9][9]bool{}
	block := [9][9]bool{}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				value := board[r][c] - '1'
				blockIndex := r/3*3 + c/3
				if row[r][value] || col[c][value] || block[blockIndex][value] {
					return false
				}
				row[r][value] = true
				col[c][value] = true
				block[blockIndex][value] = true
			}
		}
	}
	return true
}
