package structures

import (
	"errors"
	"gb-go-architecture/lesson-5/structures/list"
)

var ErrEmptyQueue = errors.New("Queue is empty")

type Queue struct {
	queue *list.DoubleLinkedList
}

func (q *Queue) Len() int {
	return q.queue.Len()
}

func (q *Queue) Push(value interface{}) {
	node := &list.DoubleListNode{Value: value}
	q.queue.PushTail(node)
}

func (q *Queue) Pop() (value interface{}, err error) {
	node := q.queue.PopHead()
	if node == nil {
		return nil, ErrEmptyQueue
	}

	return node.Value, nil
}

func NewQueue() *Queue {
	list := &list.DoubleLinkedList{}
	return &Queue{
		queue: list,
	}
}
