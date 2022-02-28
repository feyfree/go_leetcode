package p0117

import "github.com/feyfree/go_leetcode/structures"

type Node = structures.Node

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	// current 相当于是上层
	// dummy 和 pre 相当于是处理相对下一层的连接
	current := root
	for current != nil {
		dummy := &Node{}
		pre := dummy
		for current != nil {
			if current.Left != nil {
				// 这一步实际上是 最后current = dummy.Next 这一步的基础
				pre.Next = current.Left
				pre = pre.Next
			}
			if current.Right != nil {
				pre.Next = current.Right
				pre = pre.Next
			}
			current = current.Next
		}
		current = dummy.Next
	}
	return root
}
