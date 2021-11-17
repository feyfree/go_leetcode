package p0006

import "strings"

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	container := make([]strings.Builder, numRows)
	row := 0
	down := true
	for i := 0; i < len(s); i++ {
		container[row].WriteByte(s[i])
		if row == numRows-1 {
			down = false
		} else if row == 0 {
			down = true
		}
		if down {
			row++
		} else {
			row--
		}
	}
	var result strings.Builder
	for i := 0; i < len(container); i++ {
		result.WriteString(container[i].String())
	}
	return result.String()
}
