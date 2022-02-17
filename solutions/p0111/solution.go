package p0111

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	if leftDepth == 0 || rightDepth == 0 {
		return leftDepth + rightDepth + 1
	}
	if leftDepth > rightDepth {
		return rightDepth + 1
	}
	return leftDepth + 1
}
