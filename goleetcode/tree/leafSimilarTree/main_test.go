package tree

import "testing"

func TestLeafSimilar(t *testing.T) {
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}

	root2 := &TreeNode{Val: 2}
	root2.Left = &TreeNode{Val: 3}
	root2.Right = &TreeNode{Val: 2}
	if leafSimilar(root1, root2) {
		t.Errorf("error")
	}

}
