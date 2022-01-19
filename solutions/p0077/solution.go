package p0077

func combine(n int, k int) [][]int {
	var result [][]int
	var current []int
	// 起点
	s := 0
	dfs(s, 0, n, k, &result, current)
	return result
}

func dfs(s int, d int, n int, k int, result *[][]int, current []int) {
	if d == k {
		*result = append(*result, append([]int(nil), current...))
		return
	}
	for i := s + 1; i <= n; i++ {
		dfs(i, d+1, n, k, result, append(current, i))
	}
}
