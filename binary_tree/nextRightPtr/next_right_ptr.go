package main

import "fmt"

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Queue struct {
	data   [][]*Node
	length int
}

func (s *Queue) Pop() []*Node {
	dat := s.data[0]
	s.data = s.data[1:]
	s.length--
	return dat
}

func (s *Queue) Push(d []*Node) {
	s.data = append(s.data, d)
	s.length++
}

func levelOrderTraversal(root *Node) {
	q := &Queue{
		data:   make([][]*Node, 0),
		length: 0,
	}

	q.Push([]*Node{root})
	for q.length != 0 {
		e := q.Pop()
		temp := make([]*Node, 0)
		for i := range len(e) - 1 {
			e[i].Next = e[i+1]
			if e[i].Left != nil {
				temp = append(temp, e[i].Left)
			}

			if e[i].Right != nil {
				temp = append(temp, e[i].Right)
			}
		}
		e[len(e)-1].Next = nil
		if e[len(e)-1].Left != nil {
			temp = append(temp, e[len(e)-1].Left)
		}

		if e[len(e)-1].Right != nil {
			temp = append(temp, e[len(e)-1].Right)
		}
		if len(temp) != 0 {
			q.Push(temp)
		}

	}
}

func connect(root *Node) *Node {
	levelOrderTraversal(root)
	return root
}

func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
		},
		Right: &Node{
			Val: 3,
		},
	}

	connect(root)

	q := &Queue{
		data: make([][]*Node, 0),
	}

	q.Push([]*Node{root})
	for q.length != 0 {
		e := q.Pop()
		fmt.Println(e[0].Val, e[0].Next)
		if e[0].Left != nil {
			q.Push([]*Node{e[0].Left})
		}

		if e[0].Right != nil {
			q.Push([]*Node{e[0].Right})
		}
	}
}
