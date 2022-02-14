package p0109

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode
type ListNode = structures.ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	var data []int
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}
	return buildTree0(data, 0, len(data)-1)

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
