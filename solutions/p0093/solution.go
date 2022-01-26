package p0093

import (
	"strconv"
)

func restoreIpAddresses(s string) []string {
	var result []string
	dfs(s, 0, "", &result)
	return result
}

func dfs(s string, d int, ip string, result *[]string) {
	l := len(s)
	if d == 4 {
		if l == 0 {
			*result = append(*result, ip)
		}
		return
	}
	if l > 0 {
		end := 3
		if s[0] == '0' {
			end = 1
		} else {
			if l < 3 {
				end = l
			}
		}
		for i := 1; i <= end; i++ {
			ss := s[:i]
			if i == 3 {
				if val, err := strconv.Atoi(ss); err == nil {
					if val > 255 {
						return
					}
				}
			}
			if d == 0 {
				dfs(s[i:], d+1, ip+ss, result)
			} else {
				dfs(s[i:], d+1, ip+"."+ss, result)
			}
		}
	}
}
