package level

import "github.com/feyfree/go_leetcode/structures"

type Node = structures.Node

func connect(root *Node) *Node {
	order := levelOrder(root)
	for i := 0; i < len(order); i++ {
		for j := 0; j < len(order[i]); j++ {
			if j < len(order[i])-1 {
				order[i][j].Next = order[i][j+1]
			} else {
				order[i][j].Next = nil
			}
		}
	}
	return root
}

func levelOrder(root *Node) [][]*Node {
	var result [][]*Node
	dfs(root, 0, &result)
	return result
}

func dfs(root *Node, d int, result *[][]*Node) {
	if root == nil {
		return
	}
	for len(*result) <= d {
		*result = append(*result, []*Node{})
	}
	(*result)[d] = append((*result)[d], root)
	dfs(root.Left, d+1, result)
	dfs(root.Right, d+1, result)
}
