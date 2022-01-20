package p0084

import "container/list"

func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := list.New()
	result := 0
	i := 0
	for i <= n {
		var height int
		if i == n {
			height = 0
		} else {
			height = heights[i]
		}
		if stack.Len() == 0 || height >= heights[stack.Back().Value.(int)] {
			stack.PushBack(i)
			i++
		} else {
			h := heights[stack.Back().Value.(int)]
			stack.Remove(stack.Back())
			var w int
			if stack.Len() > 0 {
				w = i - stack.Back().Value.(int) - 1
			} else {
				w = i
			}
			if h*w > result {
				result = h * w
			}
		}
	}
	return result
}
