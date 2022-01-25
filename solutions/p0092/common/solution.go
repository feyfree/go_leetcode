package common

import (
	"github.com/feyfree/go_leetcode/structures"
)

type ListNode = structures.ListNode

func reverseBetween(head *ListNode, left, right int) *ListNode {
	// 设置 dummy 是这一类问题的一般做法
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}
