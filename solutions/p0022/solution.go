package p0022

func generateParenthesis(n int) []string {
	if n == 0 {
		return make([]string, 0)
	}
	var result []string
	dfs(n, n, n, &result, "")
	return result
}

func dfs(left int, right int, n int, result *[]string, current string) {
	if left == 0 && right == 0 {
		*result = append(*result, current)
		return
	}
	if left != 0 {
		dfs(left-1, right, n, result, current+"(")
	}
	if right > left {
		dfs(left, right-1, n, result, current+")")
	}
}
