package helper

type Stack struct {
	data   []int
	length int
}

func GetNewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Length() int {
	return s.length
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
	s.length++
}

func (s *Stack) Pop() int {
	ret := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.length--
	return ret
}
