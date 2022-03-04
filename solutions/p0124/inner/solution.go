package inner

import (
	"github.com/feyfree/go_leetcode/structures"
	"math"
)

type TreeNode = structures.TreeNode

func maxPathSum(root *TreeNode) int {
	var result = math.MinInt32
	if root == nil {
		return 0
	}
	var findMaxPathSum func(node *TreeNode) int

	findMaxPathSum = func(node *TreeNode) int {
		if node == nil {
			return math.MinInt32
		}
		l := max(0, findMaxPathSum(node.Left))
		r := max(0, findMaxPathSum(node.Right))
		sum := l + r + node.Val
		result = max(result, sum)
		return node.Val + max(l, r)
	}
	findMaxPathSum(root)
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
