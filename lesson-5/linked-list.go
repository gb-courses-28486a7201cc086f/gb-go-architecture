package main

type Node struct {
	next *Node
	Data int
}

type List struct {
	len  int
	head *Node
	tail *Node
}

func (l *List) Head() *Node {
	return l.head
}

func (l *List) Tail() *Node {
	return l.tail
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Find(elem int) *Node {
	if l.head != nil {
		for tmp := l.head; tmp != l.tail; tmp = tmp.next {
			if tmp.Data == elem {
				return tmp
			}
		}
	}
	return nil
}

func (l *List) Add(prev *Node, node *Node) {
	l.len++
	if prev == nil {
		node.next = l.head
		l.head = node
		return
	}
	if l.head == nil {
		l.head = node
		l.tail = l.head
		return
	}
	node.next = prev.next
	prev.next = node
	if prev == l.tail {
		l.tail = node
	}
}

func (l *List) Append(node *Node) {
	l.Add(l.tail, node)
}

func (l *List) Preppend(node *Node) {
	l.Add(nil, node)
}

func (l *List) Delete(node *Node) {
	l.len--
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	}
	if l.head != nil {
		for tmp := l.head; tmp != l.tail; tmp = tmp.next {
			if tmp.next == node && node != l.tail {
				tmp.next = node.next
			}
			if tmp.next == node && node == l.tail {
				tmp.next = nil
				l.tail = tmp
			}
		}
	}
	if node == l.head {
		l.head = node.next
	}
}
