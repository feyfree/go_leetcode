package recursion

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var queue []*TreeNode
	queue = append(queue, root)
	levelOrder0(queue, &result, 0)
	for i := 0; i < len(result); i += 2 {
		reverse(result[i])
	}
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

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
