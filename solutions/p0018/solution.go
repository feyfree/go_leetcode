package p0018

import "sort"

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return make([][]int, 0)
	}
	// 数据排序
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			start := j + 1
			end := n - 1
			for start < end {
				sum := nums[i] + nums[j] + nums[start] + nums[end] - target
				if start-1 > j && nums[start] == nums[start-1] {
					start++
					continue
				}
				if sum == 0 {
					result = append(result, []int{nums[i], nums[j], nums[start], nums[end]})
					start++
				} else if sum > 0 {
					end--
				} else {
					start++
				}
			}
		}
	}
	return result
}
