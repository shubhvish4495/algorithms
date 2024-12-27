package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal2(root *TreeNode) []int {
	result := make([]int, 0)
	var inorder func(*TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}

	inorder(root)
	return result
}

func inorderTraversal(root *TreeNode) []int {
	var traversalArr []int
	if root == nil {
		return []int{}
	}

	val := root.Val
	leftT := inorderTraversal(root.Left)
	righT := inorderTraversal(root.Right)
	traversalArr = append(traversalArr, leftT...)
	traversalArr = append(traversalArr, val)
	traversalArr = append(traversalArr, righT...)

	return traversalArr
}

func main() {
	r := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(inorderTraversal(r))

}
