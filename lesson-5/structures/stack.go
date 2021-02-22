package structures

import (
	"errors"
	"gb-go-architecture/lesson-5/structures/list"
)

var ErrEmptyStack = errors.New("Stack is empty")

type Stack struct {
	stack *list.DoubleLinkedList
}

func (s *Stack) Len() int {
	return s.stack.Len()
}

func (s *Stack) Push(value interface{}) {
	node := &list.DoubleListNode{Value: value}
	s.stack.PushHead(node)
}

func (s *Stack) Pop() (value interface{}, err error) {
	node := s.stack.PopHead()
	if node == nil {
		return nil, ErrEmptyStack
	}

	return node.Value, nil
}

func NewStack() *Stack {
	list := &list.DoubleLinkedList{}
	return &Stack{
		stack: list,
	}
}
