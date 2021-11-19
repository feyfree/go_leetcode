package p0016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < n; i++ {
		start := i + 1
		end := n - 1
		for start < end {
			sum := nums[start] + nums[end] + nums[i]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-ans)) {
				ans = sum
			}
			if sum > target {
				end--
			} else if sum < target {
				start++
			} else {
				return ans
			}
		}
	}
	return ans
}
