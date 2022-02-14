package p0106

import "github.com/feyfree/go_leetcode/structures"

type TreeNode = structures.TreeNode

func buildTree(inorder []int, postorder []int) *TreeNode {
	position := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		position[inorder[i]] = i
	}
	return buildTree0(postorder, 0, 0, len(postorder)-1, &position)
}

func buildTree0(postorder []int, iStart int, pStart int, pEnd int, position *map[int]int) *TreeNode {
	if pStart > pEnd {
		return nil
	}
	root := &TreeNode{Val: postorder[pEnd]}
	// 相当于  root 在 inorder 里面的顺序
	rootIndex := (*position)[postorder[pEnd]]
	pMid := pStart + rootIndex - iStart - 1
	root.Left = buildTree0(postorder, iStart, pStart, pMid, position)
	root.Right = buildTree0(postorder, rootIndex+1, pMid+1, pEnd-1, position)
	return root
}
