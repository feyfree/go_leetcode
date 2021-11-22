package p0025

import "github.com/feyfree/go_leetcode/structures"

type ListNode = structures.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 反转链表, 移动cursor
	if head == nil || k <= 1 {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	count := 0
	for true {
		loop := pre
		for loop != nil {
			loop = loop.Next
			count++
		}
		if count <= k {
			break
		}
		cur := pre.Next
		nxt := cur.Next
		// 这个循环里面cur是没有  cur = xxx 这种标记的, 因为在交换节点中 实际上cur已经后移动了
		for i := 1; i < k && nxt != nil; i++ {
			cur.Next = nxt.Next
			nxt.Next = pre.Next
			pre.Next = nxt
			nxt = cur.Next
		}
		pre = cur
		count = 0
	}
	return dummy.Next
}
