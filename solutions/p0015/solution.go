package p0015

import "sort"

// 数组内部可能重复
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	// 双指针
	var result [][]int
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start := i + 1
		end := n - 1
		for start < end {
			sum := nums[start] + nums[end] + nums[i]
			if start-1 > i && nums[start] == nums[start-1] {
				start++
				continue
			}
			if sum == 0 {
				result = append(result, []int{nums[i], nums[start], nums[end]})
				start++
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}
