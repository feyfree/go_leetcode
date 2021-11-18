package p0015

import "sort"

// 数组内部可能重复
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	// 二分搜索
	for i := 0; i < len(nums); i++ {
		va := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if (j-1 > i) && nums[j] == nums[j-1] {
				continue
			}
			vb := nums[j]
			start := j
			end := len(nums) - 1
			for start < end {
				mid := start + (end - start) / 2
				sum = va + vb + nums[]
			}
		}
	}

}
