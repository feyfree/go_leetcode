package p0100

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	same := (p == nil && q == nil) || (p != nil && q != nil && p.Val == q.Val)
	if same && p != nil {
		same = isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return same
}
