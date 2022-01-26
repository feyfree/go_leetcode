package p0098

import (
	"github.com/feyfree/go_leetcode/structures"
	"math"
)

type TreeNode = structures.TreeNode

var prev int

func isValidBST(root *TreeNode) bool {
	prev = math.MinInt64
	return inOrder(root)
}

func inOrder(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !inOrder(root.Left) {
		return false
	}
	if root.Val <= prev {
		return false
	}
	prev = root.Val
	return inOrder(root.Right)
}
