package p0056

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// intervals 左边界排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var result [][]int
	for _, interval := range intervals {
		if len(result) == 0 || interval[0] > result[len(result)-1][1] {
			result = append(result, interval)
		} else {
			result[len(result)-1][1] = int(math.Max(float64(result[len(result)-1][1]), float64(interval[1])))
		}
	}
	return result
}
