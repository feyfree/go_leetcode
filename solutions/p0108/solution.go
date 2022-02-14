package p0108

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	return buildTree0(nums, 0, len(nums)-1)

}

func buildTree0(nums []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildTree0(nums, start, mid-1)
	root.Right = buildTree0(nums, mid+1, end)
	return root
}
