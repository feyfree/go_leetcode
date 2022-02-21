package p0114

import (
	"github.com/feyfree/go_leetcode/structures"
)

type TreeNode = structures.TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	right := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	flatten(right)
	root.Right = right
}
