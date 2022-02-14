package p0110

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode

var balanced bool

func isBalanced(root *TreeNode) bool {
	balanced = true
	if root == nil {
		return true
	}
	height(root)
	return balanced
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := height(root.Left)
	rh := height(root.Right)

	if lh != rh && (lh-rh > 1 || rh-lh > 1) {
		balanced = false
		return -1
	}
	if lh > rh {
		return lh + 1
	} else {
		return rh + 1
	}
}
