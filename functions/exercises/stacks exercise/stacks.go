package main

import "fmt"

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	s.i--
	ret := s.data[s.i]
	s.data[s.i] = 0
	return ret
}

func main() {
	var s stack
	s.push(25)
	fmt.Printf("stack %v\n", s)
	s.push(14)
	fmt.Printf("stack %v\n", s)
}
