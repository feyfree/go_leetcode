package p0066

func plusOne(digits []int) []int {
	flag := 0
	for i := 0; i < len(digits); i++ {
		digit := digits[len(digits)-1-i]
		if i == 0 {
			digit++
		}
		digits[len(digits)-1-i] = (digit + flag) % 10
		flag = (digit + flag) / 10
	}
	if flag == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}
