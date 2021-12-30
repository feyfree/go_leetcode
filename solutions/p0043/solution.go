package p0043

func multiply(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	data := make([]byte, l1+l2)
	for i := 0; i < len(data); i++ {
		data[i] = '0'
	}
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			sum := data[i+j+1] - '0' + (num1[i]-'0')*(num2[j]-'0')
			data[i+j+1] = sum%10 + '0'
			data[i+j] += sum / 10
		}
	}
	for i := 0; i < len(data); i++ {
		if data[i] != '0' || i == len(data)-1 {
			return string(data[i:])
		}
	}
	return ""
}
