package helper

type Queue struct {
	data   []int
	length int
}

func GetNewQueue() *Queue {
	return &Queue{
		data: make([]int, 0),
	}
}

func (s *Queue) Length() int {
	return s.length
}

func (s *Queue) Push(x int) {
	s.data = append(s.data, x)
	s.length++
}

func (s *Queue) Pop() int {
	ret := s.data[0]
	s.data = s.data[1:]
	s.length--
	return ret
}
