package p0102

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var queue []*TreeNode
	queue = append(queue, root)
	levelOrder0(queue, &result, 0)
	return result
}

func levelOrder0(q []*TreeNode, result *[][]int, d int) {
	if q != nil && len(q) > 0 {
		var nextQ []*TreeNode
		*result = append(*result, []int{})
		for i := 0; i < len(q); i++ {
			(*result)[d] = append((*result)[d], q[i].Val)
			if q[i].Left != nil {
				nextQ = append(nextQ, q[i].Left)
			}
			if q[i].Right != nil {
				nextQ = append(nextQ, q[i].Right)
			}
		}
		d++
		levelOrder0(nextQ, result, d)
	}
}
