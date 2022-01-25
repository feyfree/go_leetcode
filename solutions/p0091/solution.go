package p0091

// 斐波那契数列
func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}
	if n == 1 {
		return 1
	}
	// 滚动数组解法
	a := 1
	b := 1
	for i := 1; i < n; i++ {
		c := 0
		if !isValid(s[i]) && !isValidDual(s[i-1], s[i]) {
			return 0
		}
		if isValid(s[i]) {
			c = b
		}
		if isValidDual(s[i-1], s[i]) {
			c = c + a
		}
		a = b
		b = c

	}
	return b

}

func isValid(a byte) bool {
	return a != '0'
}

func isValidDual(a byte, b byte) bool {
	num := (a-'0')*10 + (b - '0')
	return num >= 10 && num <= 26
}
