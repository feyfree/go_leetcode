package p0067

import "strings"

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		return addBinary(b, a)
	}
	container := &strings.Builder{}
	flag := 0
	i := 0
	for i < len(b) {
		ac := a[len(a)-1-i]
		bc := b[len(b)-1-i]
		sum := int(ac-'0'+bc-'0') + flag
		flag = sum / 2
		container.WriteByte(byte(sum%2 + '0'))
		i++
	}
	for i < len(a) {
		ac := a[len(a)-1-i]
		sum := int(ac-'0') + flag
		flag = sum / 2
		container.WriteByte(byte(sum%2 + '0'))
		i++
	}
	if flag == 1 {
		container.WriteByte('1')
	}
	return reverseString(container.String())

}

func reverseString(s string) string {
	bytes := []byte(s)
	for from, to := 0, len(bytes)-1; from < to; from, to = from+1, to-1 {
		bytes[from], bytes[to] = bytes[to], bytes[from]
	}
	return string(bytes)
}
