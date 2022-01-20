package p0082

import (
	"github.com/feyfree/go_leetcode/structures"
)

type ListNode = structures.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cursor := head
	previous := dummy
	for cursor != nil {
		num := 0
		for cursor.Next != nil && cursor.Next.Val == cursor.Val {
			cursor = cursor.Next
			num++
		}
		if num == 0 {
			previous.Next = cursor
			previous = previous.Next
		} else {
			previous.Next = nil
		}
		cursor = cursor.Next
	}
	return dummy.Next
}
