package p0061

import "github.com/feyfree/go_leetcode/structures"

type ListNode = structures.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	cursor := &ListNode{}
	cursor.Next = head
	total := 0
	for cursor.Next != nil {
		total += 1
		cursor = cursor.Next
	}
	k %= total
	if k == 0 {
		return head
	}
	k = total - k
	pre := head
	visit := head
	for i := 0; i < k; i++ {
		pre = visit
		visit = visit.Next
	}
	pre.Next = nil
	result := visit
	for visit.Next != nil {
		visit = visit.Next
	}
	visit.Next = head
	return result
}
