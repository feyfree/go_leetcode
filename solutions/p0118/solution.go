package p0118

func generate(numRows int) [][]int {
	var result [][]int
	result = append(result, []int{1})
	for i := 1; i < numRows; i++ {
		before := result[i-1]
		var current []int
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				current = append(current, 1)
			} else {
				current = append(current, before[j-1]+before[j])
			}
		}
		result = append(result, current)
	}
	return result
}
