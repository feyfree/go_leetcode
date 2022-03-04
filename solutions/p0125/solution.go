package p0125

import "unicode"

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left <= right {
		if !validChar(s[left]) {
			left++
			continue
		}
		if !validChar(s[right]) {
			right--
			continue
		}
		if unicode.ToUpper(rune(s[left])) != unicode.ToUpper(rune(s[right])) {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}

func validChar(a byte) bool {
	return unicode.IsDigit(rune(a)) || unicode.IsLetter(rune(a))
}
