package main

import "fmt"

type Queue struct {
	list *List
}

func NewQueue(cap int) *Queue {
	return &Queue{
		list: &List{},
	}
}

func (s *Queue) Push(elem int) {
	node := &Node{
		Data: elem,
	}
	s.list.Append(node)
}
func (s *Queue) Pop() int {
	if s.list.Len() == 0 {
		return 0
	}
	elem := s.list.Head().Data
	s.list.Delete(s.list.Head())
	return elem
}

func main() {
	queue := NewQueue(2)
	queue.Push(5)
	queue.Push(6)
	queue.Push(7)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}
