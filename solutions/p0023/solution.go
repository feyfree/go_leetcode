package p0023

import "github.com/feyfree/go_leetcode/structures"

type ListNode = structures.ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	n := len(lists)
	if n == 1 {
		return lists[0]
	} else if n == 2 {
		return mergeTwoList(lists[0], lists[1])
	} else {
		mid := n / 2
		partA := lists[:mid]
		partB := lists[mid:]
		dataA := mergeKLists(partA)
		dataB := mergeKLists(partB)
		combination := []*ListNode{dataA, dataB}
		return mergeKLists(combination)
	}
}

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
			temp = temp.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
			temp = temp.Next
		}
	}
	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}
	return result.Next
}
