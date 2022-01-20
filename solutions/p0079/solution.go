package p0079

var rows int
var cols int
var directions [5]int

func exist(board [][]byte, word string) bool {
	directions = [...]int{-1, 0, 1, 0, -1}
	rows = len(board)
	cols = len(board[0])
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if dfs(&board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board *[][]byte, word string, d int, x int, y int) bool {
	if x < 0 || x == cols || y < 0 || y == rows || word[d] != (*board)[y][x] {
		return false
	}
	if d == len(word)-1 {
		return true
	}
	current := (*board)[y][x]
	(*board)[y][x] = '0'
	found := false
	for i := 0; i < 4; i++ {
		found = found || dfs(board, word, d+1, x+directions[i], y+directions[i+1])
	}
	(*board)[y][x] = current
	return found
}
