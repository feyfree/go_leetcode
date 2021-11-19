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
	// 为什么使用fast.Next    slow -> fast fast最后一个绝对不能是nil
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
