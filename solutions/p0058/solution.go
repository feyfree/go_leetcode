package p0058

func lengthOfLastWord(s string) int {
	end := len(s) - 1
	for s[end] == ' ' {
		end--
	}
	count := 0
	for end >= 0 && s[end] != ' ' {
		count++
		end--
	}
	return count
}
