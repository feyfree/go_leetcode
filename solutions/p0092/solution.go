package p0092

import (
	"github.com/feyfree/go_leetcode/structures"
)

type ListNode = structures.ListNode

// 1 <= left <= right <= n
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	// 前驱节点
	front := dummy
	current := &ListNode{}
	prev := &ListNode{}
	tail := &ListNode{}
	for i := 0; i < left-1; i++ {
		front = front.Next
	}
	prev = front
	current = front.Next
	tail = current
	for i := left; i <= right; i++ {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	front.Next = prev
	tail.Next = current
	return dummy.Next
}
