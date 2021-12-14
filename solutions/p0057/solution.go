package p0057

import "math"

func insert(intervals [][]int, newInterval []int) [][]int {
	var data [][]int
	var result [][]int
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] >= newInterval[0] {
			data = append(data, newInterval)
		}
		data = append(data, intervals[i])
	}
	if len(data) == len(intervals) {
		data = append(data, newInterval)
	}
	for _, interval := range data {
		if len(result) == 0 || interval[0] > result[len(result)-1][1] {
			result = append(result, interval)
		} else {
			result[len(result)-1][1] = int(math.Max(float64(result[len(result)-1][1]), float64(interval[1])))
		}
	}
	return result
}
