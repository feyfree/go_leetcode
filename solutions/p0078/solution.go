package p0078

func subsets(nums []int) [][]int {
	var result [][]int
	var current []int
	// 数组长度
	n := len(nums)
	for k := 0; k <= n; k++ {
		dfs(-1, 0, n, k, &result, current, nums)
	}
	return result
}

func dfs(s int, d int, n int, k int, result *[][]int, current []int, nums []int) {
	if d == k {
		*result = append(*result, append([]int(nil), current...))
		return
	}
	for i := s + 1; i < n; i++ {
		dfs(i, d+1, n, k, result, append(current, nums[i]), nums)
	}
}
