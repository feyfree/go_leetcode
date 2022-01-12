package p0061

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	a.Next = b
	b.Next = c
	result := rotateRight(a, 2)
	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
}
