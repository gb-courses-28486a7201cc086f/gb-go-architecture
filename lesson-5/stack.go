package main

import "fmt"

type Stack struct {
	sli []int
}

func NewStack(cap int) *Stack {
	return &Stack{
		sli: make([]int, 0, cap),
	}
}

func (s *Stack) Push(elem int) {
	s.sli = append(s.sli, elem)
}
func (s *Stack) Pop() int {
	if len(s.sli) == 0 {
		return 0
	}
	elem := s.sli[len(s.sli)-1]
	s.sli = s.sli[:len(s.sli)-1]
	return elem
}

func main() {
	stack := NewStack(2)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
