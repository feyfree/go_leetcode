package p0086

import "github.com/feyfree/go_leetcode/structures"

type ListNode = structures.ListNode

func partition(head *ListNode, x int) *ListNode {
	dummyL := &ListNode{}
	pl := dummyL
	dummyR := &ListNode{}
	pr := dummyR
	for head != nil {
		if head.Val < x {
			pl.Next = head
			pl = pl.Next
		} else {
			pr.Next = head
			pr = pr.Next
		}
		head = head.Next
	}
	pr.Next = nil
	pl.Next = dummyR.Next
	return dummyL.Next
}
