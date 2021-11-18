package p0011

import "math"

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	start := 0
	end := len(height) - 1
	result := 0
	for start < end {
		if height[end] > height[start] {
			result = int(math.Max(float64((end-start)*height[start]), float64(result)))
			start++
		} else {
			result = int(math.Max(float64((end-start)*height[end]), float64(result)))
			end--
		}
	}
	return result
}
