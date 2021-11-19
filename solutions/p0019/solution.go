package p0019

import "github.com/feyfree/go_leetcode/structures"

type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	var fast = dummy
	for n > 0 {
		fast = fast.Next
		n--
	}
	var slow = dummy
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
