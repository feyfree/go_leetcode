package p0014

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	// find shortest
	shortest := 0
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < len(strs[shortest]) {
			shortest = i
		}
	}
	var result strings.Builder
	for i := 0; i < len(strs[shortest]); i++ {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != strs[shortest][i] {
				return result.String()
			}
		}
		result.WriteByte(strs[shortest][i])
	}
	return result.String()
}
