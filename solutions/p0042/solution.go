package p0042

import "math"

func trap(height []int) int {
	n := len(height)
	pre := make([]int, n)
	after := make([]int, n)
	var ans int
	for i := 0; i < n; i++ {
		if i == 0 {
			pre[i] = height[i]
		} else {
			pre[i] = int(math.Max(float64(pre[i-1]), float64(height[i])))
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			after[i] = height[i]
		} else {
			after[i] = int(math.Max(float64(after[i+1]), float64(height[i])))
		}
	}
	for i := 0; i < n; i++ {
		ans += int(math.Min(float64(pre[i]), float64(after[i]))) - height[i]
	}
	return ans
}
