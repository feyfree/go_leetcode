package recursion

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode

func isSymmetric(root *TreeNode) bool {
	return symmetric0(root.Left, root.Right)
}

func symmetric0(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}
	if a.Val == b.Val {
		return symmetric0(a.Left, b.Right) && symmetric0(a.Right, b.Left)
	}
	return false
}
