package p0060

import "strings"

func getPermutation(n int, k int) string {
	num := make([]int, 10)
	fact := make([]int, 10)
	for i := 0; i < len(fact); i++ {
		fact[i] = 1
	}
	for i := 1; i <= 9; i++ {
		num[i-1] = i
		fact[i] = fact[i-1] * i
	}
	s := strings.Builder{}
	k--
	for n > 0 {
		n--
		// 查找阶乘所在数
		d := k / fact[n]
		// 查找偏移量
		k %= fact[n]
		s.WriteByte(byte('0' + num[d]))
		for i := d + 1; i <= 9; i++ {
			num[i-1] = num[i]
		}
	}
	return s.String()
}
