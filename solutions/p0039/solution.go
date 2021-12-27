package p0039

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var current []int
	// 排序
	sort.Ints(candidates)
	min := candidates[0]
	if min > target {
		return result
	}
	maxDepth := target / min
	// dfs 搜索
	for n := 1; n <= maxDepth; n++ {
		dfs(candidates, target, 0, 0, n, current, &result)
	}
	return result
}

func dfs(candidates []int, target int, d int, s int, n int, current []int, result *[][]int) {
	if d == n {
		if target == 0 {
			*result = append(*result, append([]int(nil), current...))
			return
		}
	}
	for i := s; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		dfs(candidates, target-candidates[i], d+1, i, n, append(current, candidates[i]), result)
	}
}
