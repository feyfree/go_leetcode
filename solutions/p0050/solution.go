package p0050

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return myPow0(x, n)
	}
	return 1.0 / myPow0(x, n)
}

func myPow0(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n&1 == 1 {
		return myPow0(x*x, n/2) * x
	}
	return myPow0(x*x, n/2)
}
