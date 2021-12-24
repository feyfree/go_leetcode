package p0037

type marked struct {
	row [9][9]bool
	col [9][9]bool
	box [9][9]bool
}

func solveSudoku(board [][]byte) {
	m := &marked{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				value := board[i][j] - '1'
				boxNo := i/3*3 + j/3
				m.row[i][value] = true
				m.col[j][value] = true
				m.box[boxNo][value] = true
			}
		}
	}
	fill(board, 0, 0, m)

}

func fill(board [][]byte, r int, c int, m *marked) bool {
	if r == 9 {
		return true
	}
	nextC := (c + 1) % 9
	var nextR int
	if nextC == 0 {
		nextR = r + 1
	} else {
		nextR = r
	}
	if board[r][c] != '.' {
		return fill(board, nextR, nextC, m)
	}
	for i := '1'; i <= '9'; i++ {
		boxNo := r/3*3 + c/3
		value := i - '1'
		if !m.row[r][value] && !m.col[c][value] && !m.box[boxNo][value] {
			m.row[r][value] = true
			m.col[c][value] = true
			m.box[boxNo][value] = true
			board[r][c] = byte(i)
			if fill(board, nextR, nextC, m) {
				return true
			}
			board[r][c] = '.'
			m.row[r][value] = false
			m.col[c][value] = false
			m.box[boxNo][value] = false
		}
	}
	return false
}
