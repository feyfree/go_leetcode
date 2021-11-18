package p0398

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (solution *Solution) Pick(target int) int {
	index := -1
	n := 0
	for i := 0; i < len(solution.nums); i++ {
		if solution.nums[i] != target {
			continue
		}
		n++
		if rand.Int()%n == 0 {
			index = i
		}
	}
	return index
}
