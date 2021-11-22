package p0024

import "github.com/feyfree/go_leetcode/structures"

type ListNode = structures.ListNode

// O -> A -> B -> C  => O -> B -> A -> C
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
		nodeA := cur.Next
		nodeB := nodeA.Next
		cur.Next = nodeB
		nodeC := nodeB.Next
		nodeB.Next = nodeA
		nodeA.Next = nodeC
		cur = cur.Next.Next
	}
	return dummy.Next
}
