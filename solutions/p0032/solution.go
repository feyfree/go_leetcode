package p0032

import (
	"container/list"
	"math"
)

func longestValidParentheses(s string) int {
	result := 0
	n := len(s)
	if n == 0 {
		return result
	}
	// 如果括号进行了消除
	start := 0
	stack := list.New()
	for i := 0; i < n; i++ {
		c := s[i]
		if c == '(' {
			stack.PushBack(i)
		} else {
			if stack.Len() == 0 {
				start = i + 1
			} else {
				element := stack.Back()
				stack.Remove(element)
				var temp int
				if stack.Len() == 0 {
					temp = i - start + 1
				} else {
					temp = i - stack.Back().Value.(int)
				}
				result = int(math.Max(float64(temp), float64(result)))
			}
		}
	}
	return result
}
