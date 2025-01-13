package main

import "github.com/shubhvish4495/algorithms/helper"

func searchBST(root *helper.TreeNode, val int) *helper.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)
}

func main() {
	r := &helper.TreeNode{
		Val: 4,
		Left: &helper.TreeNode{
			Val: 2,
			Left: &helper.TreeNode{
				Val: 1,
			},
			Right: &helper.TreeNode{
				Val: 4,
			},
		},
		Right: &helper.TreeNode{
			Val: 7,
		},
	}

	t := searchBST(r, 2)
	t.Print()
}
