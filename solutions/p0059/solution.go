package p0059

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	current := 1
	total := n * n
	times := 0
	n--
	// 表明维度
	d := 0
	// 当前坐标 x
	x := 0
	// 当前坐标 y
	y := 0

	// 4个边界
	l := 0
	t := 0
	r := n
	b := n
	for times < total {
		times += 1
		if d == 0 {
			for x < r {
				result[y][x] = current
				current++
				x++
			}
			t++
		} else if d == 1 {
			for y < b {
				result[y][x] = current
				y++
				current++
			}
			r--
		} else if d == 2 {
			for x > l {
				result[y][x] = current
				x--
				current++
			}
			b--
		} else {
			for y > t {
				result[y][x] = current
				current++
				y--
			}
			l++
		}
		d = (d + 1) % 4
	}
	result[y][x] = total
	return result
}
