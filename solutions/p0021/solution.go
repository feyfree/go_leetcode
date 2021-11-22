package p0021

import "github.com/feyfree/go_leetcode/structures"

type ListNode = structures.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	cursor := result
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cursor.Next = l2
			l2 = l2.Next
			cursor = cursor.Next
		} else {
			cursor.Next = l1
			l1 = l1.Next
			cursor = cursor.Next
		}
	}
	for l1 != nil {
		cursor.Next = l1
	}
	for l2 != nil {
		cursor.Next = l2
	}
	return result.Next
}
