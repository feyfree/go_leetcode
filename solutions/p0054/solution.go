package p0054

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return make([]int, 0)
	}
	cols := len(matrix[0])
	if cols == 0 {
		return make([]int, 0)
	}
	total := rows * cols
	var result []int
	rows--
	cols--
	left := 0
	right := cols
	top := 0
	bottom := rows
	x := 0
	y := 0
	d := 0
	for len(result) < total-1 {
		if d == 0 {
			for x < right {
				result = append(result, matrix[y][x])
				x++
			}
			top++
		} else if d == 1 {
			for y < bottom {
				result = append(result, matrix[y][x])
				y++
			}
			right--
		} else if d == 2 {
			for x > left {
				result = append(result, matrix[y][x])
				x--
			}
			bottom--
		} else {
			for y > top {
				result = append(result, matrix[y][x])
				y--
			}
			left++
		}
		d = (d + 1) % 4
	}
	if len(result) != total {
		result = append(result, matrix[y][x])
	}
	return result
}
