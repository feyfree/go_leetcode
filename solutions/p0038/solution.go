package p0038

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	data := "1"
	for n-1 > 0 {
		data = generateNext(data)
		n--
	}
	return data
}

func generateNext(data string) string {
	result := strings.Builder{}
	var before byte
	count := 0
	for i := 0; i < len(data); i++ {
		if before == 0 || before == data[i] {
			count += 1
		} else {
			result.WriteString(strconv.Itoa(count))
			result.WriteByte(before)
			count = 1
		}
		before = data[i]
	}
	if count > 0 {
		result.WriteString(strconv.Itoa(count))
		result.WriteByte(before)
	}
	return result.String()
}
