package p0053

import "math"

func maxSubArray(nums []int) int {
	belowZero := true
	m := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			belowZero = false
			break
		}
		m = int(math.Max(float64(m), float64(nums[i])))
	}
	if belowZero {
		return m
	}
	maxSoFar := 0
	maxEndingHere := 0
	for _, num := range nums {
		maxEndingHere = int(math.Max(float64(maxEndingHere+num), 0))
		maxSoFar = int(math.Max(float64(maxSoFar), float64(maxEndingHere)))
	}
	return maxSoFar
}
