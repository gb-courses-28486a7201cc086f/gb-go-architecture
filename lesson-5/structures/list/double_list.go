package list

import "errors"

var ErrNotExists = errors.New("Element of list does not exists")

// DoubleListNode is element of list
type DoubleListNode struct {
	Value interface{}
	next  *DoubleListNode
	prev  *DoubleListNode
	list  *DoubleLinkedList
}

// Next iterates from node till head of list.
// Returns nil when head has been riched
func (n *DoubleListNode) Next() *DoubleListNode {
	return n.next
}

// Prev iterates from node till tail of list.
// Returns nil when tail has been riched
func (n *DoubleListNode) Prev() *DoubleListNode {
	return n.prev
}

// DoubleLinkedList provides methods for working with double linked list
type DoubleLinkedList struct {
	len  int
	head *DoubleListNode
	tail *DoubleListNode
}

// linkNodes connects two nodes to each other and to list instance
func (l *DoubleLinkedList) linkNodes(prevNode, nextNode *DoubleListNode) {
	if prevNode != nil {
		prevNode.next = nextNode
		prevNode.list = l
	}
	if nextNode != nil {
		nextNode.prev = prevNode
		nextNode.list = l
	}
}

// linkNodes disconnects two nodes from each other and from list instance
func (l *DoubleLinkedList) unlinkNodes(prevNode, nextNode *DoubleListNode) {
	if prevNode != nil {
		prevNode.next = nil
		prevNode.list = nil
	}
	if nextNode != nil {
		nextNode.prev = nil
		nextNode.list = nil
	}
}

// Head returns first element of list
func (l *DoubleLinkedList) Head() *DoubleListNode {
	return l.head
}

// Tail returns last element of list
func (l *DoubleLinkedList) Tail() *DoubleListNode {
	return l.tail
}

// Len returns elements count
func (l *DoubleLinkedList) Len() int {
	return l.len
}

// PushHead appends new element in front of list
func (l *DoubleLinkedList) PushHead(item *DoubleListNode) {
	if item == nil {
		return
	}

	l.len++
	l.linkNodes(l.head, item)

	l.head = item
	// tail == nil means than first item pushed =>
	// tail should point to that item too
	if l.tail == nil {
		l.tail = item
	}
}

// PushTail appends new element in back of list
func (l *DoubleLinkedList) PushTail(item *DoubleListNode) {
	if item == nil {
		return
	}

	l.len++
	l.linkNodes(item, l.tail)

	l.tail = item
	// head == nil means than first item pushed =>
	// head should point to that item too
	if l.head == nil {
		l.head = item
	}
}

// PopHead disconnects element from front of list and returns it
func (l *DoubleLinkedList) PopHead() *DoubleListNode {
	// empty list
	if l.head == nil {
		return nil
	}

	l.len--
	result := l.head
	l.head = l.head.prev

	// means we have removed last item =>
	// tail should be nil too
	if l.head == nil {
		l.tail = nil
	}

	l.unlinkNodes(l.head, result)
	return result
}

// PopTail disconnects element from front of list and returns it
func (l *DoubleLinkedList) PopTail() *DoubleListNode {
	// empty list
	if l.tail == nil {
		return nil
	}

	l.len--
	result := l.tail
	l.tail = l.tail.next

	// means we have removed last item =>
	// head should be nil too
	if l.tail == nil {
		l.head = nil
	}

	l.unlinkNodes(result, l.head)
	return result
}

// InsertBefore appends newItem before item.
// item should be pointer to element of current list instance.
func (l *DoubleLinkedList) InsertBefore(item, newItem *DoubleListNode) error {
	if item == nil || item.list != l {
		return ErrNotExists
	}

	l.len++
	l.linkNodes(item.prev, newItem)
	l.linkNodes(newItem, item)

	return nil
}

// InsertAfter appends newItem after item.
// item should be pointer to element of current list instance.
func (l *DoubleLinkedList) InsertAfter(item, newItem *DoubleListNode) error {
	if item == nil || item.list != l {
		return ErrNotExists
	}

	l.len++
	l.linkNodes(newItem, item.next)
	l.linkNodes(item, newItem)

	return nil
}

// Remove element from list
// item should be pointer to element of current list instance.
func (l *DoubleLinkedList) Remove(item *DoubleListNode) error {
	if item == nil || item.list != l {
		return ErrNotExists
	}

	l.len--
	l.linkNodes(item.prev, item.next)
	item.next = nil
	item.prev = nil

	return nil
}
