package p0045

import "math"

func jump(nums []int) int {
	near := 0
	far := 0
	step := 0
	for i := 0; i < len(nums); i++ {
		if i > near {
			step++
			near = far
		}
		far = int(math.Max(float64(far), float64(i+nums[i])))
	}
	return step
}
