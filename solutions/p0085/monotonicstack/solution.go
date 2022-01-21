package dp

import "container/list"

func maximalRectangle(matrix [][]byte) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	heights := make([]int, cols)
	result := 0
	for _, chars := range matrix {
		for i := 0; i < cols; i++ {
			if chars[i] == '1' {
				heights[i]++
			} else {
				heights[i] = 0
			}
		}
		cal := largestRectangle(heights)
		if cal > result {
			result = cal
		}
	}
	return result
}

func largestRectangle(heights []int) int {
	n := len(heights)
	stack := list.New()
	result := 0
	end := 0
	for end <= n {
		var height int
		if end == n {
			height = 0
		} else {
			height = heights[end]
		}
		if stack.Len() == 0 || height >= heights[stack.Back().Value.(int)] {
			stack.PushBack(end)
			end++
		} else {
			h := heights[stack.Back().Value.(int)]
			stack.Remove(stack.Back())
			var w int
			if stack.Len() == 0 {
				w = end
			} else {
				w = end - stack.Back().Value.(int) - 1
			}
			if h*w > result {
				result = h * w
			}
		}
	}
	return result
}
