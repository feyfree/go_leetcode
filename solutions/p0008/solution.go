package p0008

import (
	"math"
)

func myAtoi(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	index := 0
	result := 0
	sign := 1
	for index < n && s[index] == ' ' {
		index++
	}
	if index < n && (s[index] == '+' || s[index] == '-') {
		if s[index] == '-' {
			sign = -1
		}
		index++
	}
	for index < n {
		digit := s[index] - '0'
		if digit < 0 || digit > 9 {
			break
		}
		overflow := math.MaxInt32/10 < result || (math.MaxInt32/10 == result && math.MaxInt32%10 < digit)
		if overflow {
			if sign > 0 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		result = result*10 + int(digit)
		index++
	}
	return result * sign
}
