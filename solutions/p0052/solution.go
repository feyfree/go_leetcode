package p0052

type BoardInfo struct {
	n     int
	board [][]byte
	cols  []bool
	diag1 []bool
	diag2 []bool
}

var b *BoardInfo

func totalNQueens(n int) int {
	b = &BoardInfo{}
	b.n = n
	b.board = make([][]byte, n)
	for i := 0; i < len(b.board); i++ {
		b.board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			b.board[i][j] = '.'
		}
	}
	var result int
	b.cols = make([]bool, n)
	b.diag1 = make([]bool, 2*n-1)
	b.diag2 = make([]bool, 2*n-1)
	nQueens0(0, &result)
	return result
}

func available(x int, y int) bool {
	return !b.cols[x] && !b.diag1[x+y] && !b.diag2[x-y+b.n-1]
}

func updateBoard(x int, y int, isPut bool) {
	b.cols[x] = isPut
	b.diag1[x+y] = isPut
	b.diag2[x-y+b.n-1] = isPut
	if isPut {
		b.board[y][x] = 'Q'
	} else {
		b.board[y][x] = '.'
	}
}

func nQueens0(y int, result *int) {
	if y == b.n {
		*result++
		return
	}
	for x := 0; x < b.n; x++ {
		if !available(x, y) {
			continue
		}
		updateBoard(x, y, true)
		nQueens0(y+1, result)
		updateBoard(x, y, false)
	}
}
