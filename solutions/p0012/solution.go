package p0012

import "strings"

func intToRoman(num int) string {
	data := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX",
		"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC",
		"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM",
		"M", "MM", "MMM"}
	container := [4]int{}
	var result strings.Builder
	i := 3
	for num > 0 {
		container[i] = num % 10
		num = num / 10
		i--
	}
	for j := 0; j < 4; j++ {
		if container[j] != 0 {
			result.WriteString(data[9*(3-j)+container[j]-1])
		}
	}
	return result.String()
}
