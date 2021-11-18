package p0013

func romanToInt(s string) int {
	values := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	result := 0
	for i := 0; i < len(s)-1; i++ {
		current := values[s[i]]
		after := values[s[i+1]]
		if current < after {
			result -= current
		} else {
			result += current
		}
	}
	result += values[s[len(s)-1]]
	return result
}
