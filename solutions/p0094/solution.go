package p0094

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode

func inorderTraversal(root *TreeNode) []int {
	var result []int
	traverse(root, &result)
	return result
}

func traverse(node *TreeNode, result *[]int) {
	if node != nil {
		traverse(node.Left, result)
		*result = append(*result, node.Val)
		traverse(node.Right, result)
	}
}
