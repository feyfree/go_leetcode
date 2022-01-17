package p0070

func climbStairs(n int) int {
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		temp := b
		b = a + b
		a = temp
	}
	return b
}
