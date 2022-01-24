package recursion

// leetcode 上面测试用例更新, 纯递归会超时
// 需要增加记忆化
var marked = make(map[string]bool)

func isScramble(s1 string, s2 string) bool {
	key := s1 + s2
	val, exist := marked[key]
	if exist {
		return val
	}
	if s1 == s2 {
		marked[key] = true
		return true
	}
	if !sameCounts(s1, s2) {
		return false
	}
	n := len(s1)
	for i := 1; i < n; i++ {
		valid := isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) || isScramble(s1[:i], s2[n-i:n]) && isScramble(s1[i:], s2[:n-i])
		if valid {
			marked[key] = true
			return true
		}
	}
	marked[key] = false
	return false
}

func sameCounts(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	a1 := make([]int, 26)
	a2 := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		a1[s1[i]-'a']++
		a2[s2[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
