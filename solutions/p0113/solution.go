package p0113

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode

func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var container []int
	dfs(&result, container, root, targetSum)
	return result
}

func dfs(result *[][]int, container []int, current *TreeNode, targetSum int) {
	if current == nil {
		return
	}
	if current.Left == nil && current.Right == nil {
		if targetSum == current.Val {
			container = append(container, current.Val)
			*result = append(*result, append([]int(nil), container...))
		}
		return
	}
	dfs(result, append(container, current.Val), current.Left, targetSum-current.Val)
	dfs(result, append(container, current.Val), current.Right, targetSum-current.Val)
}
