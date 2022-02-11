package p0105

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	position := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		position[inorder[i]] = i
	}
	return buildTree0(preorder, 0, 0, len(preorder)-1, &position)
}

func buildTree0(preorder []int, iStart int, pStart int, pEnd int, position *map[int]int) *TreeNode {
	if pStart > pEnd {
		return nil
	}

	root := &TreeNode{Val: preorder[pStart]}
	// 相当于  root 在 inorder 里面的顺序
	rootIndex := (*position)[preorder[pStart]]
	pMid := pStart + rootIndex - iStart
	root.Left = buildTree0(preorder, iStart, pStart+1, pMid, position)
	root.Right = buildTree0(preorder, rootIndex+1, pMid+1, pEnd, position)
	return root
}
