package p0083

import "github.com/feyfree/go_leetcode/structures"

type ListNode = structures.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cursor := head
	for cursor != nil && cursor.Next != nil {
		if cursor.Val == cursor.Next.Val {
			cursor.Next = cursor.Next.Next
		} else {
			cursor = cursor.Next
		}
	}
	return dummy.Next
}
