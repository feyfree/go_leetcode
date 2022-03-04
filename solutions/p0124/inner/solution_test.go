package inner

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	root := &TreeNode{}
	root.Val = 0
	sum := maxPathSum(root)
	fmt.Println(sum)

}
