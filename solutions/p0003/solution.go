package p0003

import (
	"math"
)

// tips: last index 滑动窗口
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 上次出现的数组
	last := [256]int{}
	result := 0
	for i := 0; i < len(last); i++ {
		last[i] = -1
	}
	var start int
	for j := 0; j < len(s); j++ {
		start = int(math.Max(float64(start), float64(last[s[j]]+1)))
		result = int(math.Max(float64(result), float64(j-start+1)))
		last[s[j]] = j
	}
	return result
}
