package p0095

import (
	"github.com/feyfree/go_leetcode/structures"
)

type TreeNode = structures.TreeNode

func generateTrees(n int) []*TreeNode {
	var result []*TreeNode
	if n == 0 {
		return result
	}
	return generateTreesByInterval(1, n)
}

func generateTreesByInterval(l int, r int) []*TreeNode {
	var ans []*TreeNode
	if l > r {
		//fmt.Println(l, r)
		// 这地方为什么需要nil, 如果 i = 1 作为root, 那么left tree 只能是nil
		ans = append(ans, nil)
		return ans
	}
	for i := l; i <= r; i++ {
		// 注意这个地方是可能存在  l > i - 1 的 当 i = l 的时候
		for _, left := range generateTreesByInterval(l, i-1) {
			for _, right := range generateTreesByInterval(i+1, r) {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				ans = append(ans, root)
			}
		}
	}
	return ans
}
