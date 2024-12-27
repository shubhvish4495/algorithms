package main

import "fmt"

type Stack struct {
	data []int
}

func GetNewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() int {
	ret := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return ret
}

type Queue struct {
	s1 *Stack
	s2 *Stack
}

func (q *Queue) Push(x int) {

	for range len(q.s1.data) {
		temp := q.s1.Pop()
		q.s2.Push(temp)
	}

	q.s1.Push(x)

	for range len(q.s2.data) {
		temp := q.s2.Pop()
		q.s1.Push(temp)
	}

}

func (q *Queue) Pop() {
	ret := q.s1.Pop()
	fmt.Println(ret)
}

func GetNewQueue() *Queue {
	return &Queue{
		s1: GetNewStack(),
		s2: GetNewStack(),
	}
}

func main() {

	q := GetNewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	q.Pop()
	q.Pop()
	q.Pop()

}
