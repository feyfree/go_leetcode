package p0055

import "math"

func canJump(nums []int) bool {
	far := 0
	for i := 0; i < len(nums); i++ {
		if i <= far {
			far = int(math.Max(float64(far), float64(i+nums[i])))
		}
		if far >= len(nums)-1 {
			return true
		}
	}
	return false
}
