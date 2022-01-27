package p0099

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode

var prev *TreeNode
var first *TreeNode
var second *TreeNode

func recoverTree(root *TreeNode) {
	prev, first, second = nil, nil, nil
	inOrder(root)
	first.Val, second.Val = second.Val, first.Val
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	if prev != nil && root.Val < prev.Val {
		first = root
		if second == nil {
			second = prev
		} else {
			return
		}
	}
	prev = root
	inOrder(root.Right)
}
