package p0074

// 二分搜索
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	start := 0
	end := m*n - 1
	for start <= end {
		mid := start + (end-start)/2
		row := mid / n
		col := mid % n
		val := matrix[row][col]
		if val == target {
			return true
		} else if val > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}
