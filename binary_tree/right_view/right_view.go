package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var rightView []int

	return rightView
}

// https://leetcode.com/problems/binary-tree-right-side-view/description/
func main() {
	r := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(rightSideView(r))
}
