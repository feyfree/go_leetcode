package dfs

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	dfs(root, 0, &result)
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
