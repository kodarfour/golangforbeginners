package stacks

type stack struct {
	i    int
	data [10]int
}

func (s *stack) Push(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *stack) Pop() int {
	s.i--
	ret := s.data[s.i]
	s.data[s.i] = 0
	return ret
}
