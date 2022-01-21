package p0084

import "container/list"

func largestRectangleArea(heights []int) int {
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
			// 这个过程中 end 不会 ++  所以在这个过程中实际上是 end 确定, start 不断变化的过程
			h := heights[stack.Back().Value.(int)]
			stack.Remove(stack.Back())
			var w int
			if stack.Len() > 0 {
				w = end - stack.Back().Value.(int) - 1
			} else {
				w = end
			}
			if h*w > result {
				result = h * w
			}
		}
	}
	return result
}
