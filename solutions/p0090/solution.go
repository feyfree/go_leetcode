package p0090

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var result [][]int
	var current []int
	dfs(n, 0, current, &result, nums)
	return result
}

func dfs(n int, d int, current []int, result *[][]int, nums []int) {
	// 切片拷贝很重要
	*result = append(*result, append([]int(nil), current...))
	if len(current) == n {
		return
	}
	for i := d; i < n; i++ {
		if i > d && nums[i] == nums[i-1] {
			continue
		}
		dfs(n, i+1, append(current, nums[i]), result, nums)
	}
}
