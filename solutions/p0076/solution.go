package p0076

import (
	"math"
)

func minWindow(s string, t string) string {
	counter := make([]int, 256)
	for i := 0; i < len(t); i++ {
		counter[t[i]]++
	}
	from := 0
	total := len(t)
	min := math.MaxInt32
	start := 0
	for end := 0; end < len(s); end++ {
		// 向后遍历, 如果字符数组存在, 做减法
		if counter[s[end]] > 0 {
			total--
		}
		counter[s[end]]--
		for total == 0 {
			if end-start+1 < min {
				min = end - start + 1
				from = start
			}
			counter[s[start]]++
			if counter[s[start]] > 0 {
				total++
			}
			start++
		}
	}
	if min == math.MaxInt32 {
		return ""
	}
	return s[from : from+min]
}
