package global

import (
	"github.com/feyfree/go_leetcode/structures"
	"math"
)

type TreeNode = structures.TreeNode

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := math.MinInt
	findMaxPathSum(root, &result)
	return result
}

func findMaxPathSum(root *TreeNode, result *int) int {
	if root == nil {
		return math.MinInt32
	}
	l := max(0, findMaxPathSum(root.Left, result))
	r := max(0, findMaxPathSum(root.Right, result))
	sum := l + r + root.Val
	if *result < sum {
		*result = sum
	}
	return root.Val + max(l, r)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
