package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	data   [][]*TreeNode
	length int
}

func (s *Queue) Pop() []*TreeNode {
	dat := s.data[0]
	s.data = s.data[1:]
	s.length--
	return dat
}

func (s *Queue) Push(d []*TreeNode) {
	s.data = append(s.data, d)
	s.length++
}

func levelOrderTraversal(root *TreeNode) [][]int {
	var lvlOrderTrvsl [][]int
	q := &Queue{
		data:   make([][]*TreeNode, 0),
		length: 0,
	}

	if root != nil {
		q.Push([]*TreeNode{root})
	}

	for q.length != 0 {
		t := q.Pop()
		tempq := make([]*TreeNode, 0)
		tempTrvsl := make([]int, 0)
		for i := range t {
			tempTrvsl = append(tempTrvsl, t[i].Val)
			if t[i].Left != nil {
				tempq = append(tempq, t[i].Left)
			}
			if t[i].Right != nil {
				tempq = append(tempq, t[i].Right)
			}
		}
		if len(tempq) != 0 {
			q.Push(tempq)
		}

		lvlOrderTrvsl = append(lvlOrderTrvsl, tempTrvsl)
	}

	return lvlOrderTrvsl
}

func leftViewOfTree(root *TreeNode) []int {
	leftView := make([]int, 0)
	lvlTraversal := levelOrderTraversal(root)
	for i := range lvlTraversal {
		leftView = append(leftView, lvlTraversal[i][0])
	}
	return leftView
}

func rightSideView(root *TreeNode) []int {
	rightView := make([]int, 0)
	lvlTraversal := levelOrderTraversal(root)
	for i := range lvlTraversal {
		n := len(lvlTraversal[i])
		rightView = append(rightView, lvlTraversal[i][n-1])
	}
	return rightView
}

func main() {

	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 10,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
		},
	}
	fmt.Println(leftViewOfTree(t))
}
