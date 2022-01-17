package newton

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	r, lx := int64(x), int64(x)
	for r > lx/r {
		r = (r + lx/r) / 2
	}
	return int(r)
}
