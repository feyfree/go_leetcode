package recursion

import "github.com/feyfree/go_leetcode/structures"

type Node = structures.Node

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left == nil || root.Right == nil {
		return root
	}
	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	// 实际上 递归过程中 重要的是串联关系 而不是返回值
	// 用 _ 标识一下方便理解
	_ = connect(root.Left)
	_ = connect(root.Right)
	return root
}
