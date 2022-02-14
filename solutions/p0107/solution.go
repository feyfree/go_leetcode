package p0107

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	dfs(root, 0, &result)
	reverse(&result)
	return result
}

func dfs(root *TreeNode, d int, result *[][]int) {
	if root == nil {
		return
	}
	for len(*result) <= d {
		*result = append(*result, []int{})
	}
	(*result)[d] = append((*result)[d], root.Val)
	dfs(root.Left, d+1, result)
	dfs(root.Right, d+1, result)
}

func reverse(s *[][]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
