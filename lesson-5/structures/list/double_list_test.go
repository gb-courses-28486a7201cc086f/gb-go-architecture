package list

import "testing"

var (
	testNodes []*DoubleListNode = []*DoubleListNode{
		{Value: 1},
		{Value: 2},
		{Value: 3},
	}
	testNodeToRemove = testNodes[1]
	testOtherNode    = &DoubleListNode{Value: 100}
)

func TestEmpty(t *testing.T) {
	list := DoubleLinkedList{}

	expectedLen := 0
	var expectedItem *DoubleListNode

	t.Run("new empty list", func(t *testing.T) {
		if list.Len() != 0 {
			t.Errorf("invalid len, got %d, expected %d", list.Len(), expectedLen)
		}
		if list.Head() != nil {
			t.Errorf("invalid head, got %v, expected %v", list.Head(), expectedItem)
		}
		if list.Tail() != nil {
			t.Errorf("invalid tail, got %v, expected %v", list.Tail(), expectedItem)
		}
	})

	t.Run("push nil", func(t *testing.T) {
		list.PushHead(nil)
		list.PushTail(nil)

		if list.Len() != 0 {
			t.Errorf("invalid len, got %d, expected %d", list.Len(), expectedLen)
		}
		if list.Head() != nil {
			t.Errorf("invalid head, got %v, expected %v", list.Head(), expectedItem)
		}
		if list.Tail() != nil {
			t.Errorf("invalid tail, got %v, expected %v", list.Tail(), expectedItem)
		}
	})

	t.Run("pop from empty", func(t *testing.T) {
		head := list.PopHead()
		tail := list.PopTail()

		if head != nil {
			t.Errorf("invalid head, got %v, expected %v", head, expectedItem)
		}
		if tail != nil {
			t.Errorf("invalid tail, got %v, expected %v", tail, expectedItem)
		}

	})
}

func TestHead(t *testing.T) {
	list := DoubleLinkedList{}

	t.Run("push values", func(t *testing.T) {
		for i, node := range testNodes {
			list.PushHead(node)

			newLen := list.Len()
			expectedLen := i + 1

			if newLen != expectedLen {
				t.Errorf("invalid list lenght, got %d, expected %d", newLen, expectedLen)
			}

			head := list.Head()
			if head != node {
				t.Errorf("invalid head, got %v, expected %v", head, node)
			}
		}
	})

	t.Run("iterate from head", func(t *testing.T) {
		item := list.Head()
		for i := list.Len() - 1; i >= 0; i-- {
			expectedItem := testNodes[i]

			if item != expectedItem {
				t.Errorf("invalid item, got %v, expedted, %v", item, expectedItem)
			}

			item = item.Prev()
		}

		// last value should be nil to stop iteration
		if item != nil {
			t.Errorf("invalid item, got %v, expedted, %v", item, nil)
		}
	})

	t.Run("pop values", func(t *testing.T) {
		for i := len(testNodes) - 1; i >= 0; i-- {
			expedtedNode := testNodes[i]
			expectedLen := i

			node := list.PopHead()
			newLen := list.Len()

			if newLen != expectedLen {
				t.Errorf("invalid list lenght, got %d, expected %d", newLen, expectedLen)
			}

			if node != expedtedNode {
				t.Errorf("invalid head, got %v, expected %v", node, expedtedNode)
			}
		}

		// we "pop" all items, expedcted empty list
		if list.Len() != 0 {
			t.Errorf("invalid len, got %d, expected %d", list.Len(), 0)
		}
		if list.Head() != nil {
			t.Errorf("invalid head, got %v, expected %v", list.Head(), nil)
		}
		if list.Tail() != nil {
			t.Errorf("invalid tail, got %v, expected %v", list.Tail(), nil)
		}
	})
}

func TestTail(t *testing.T) {
	list := DoubleLinkedList{}

	t.Run("push values", func(t *testing.T) {
		for i, node := range testNodes {
			list.PushTail(node)

			newLen := list.Len()
			expectedLen := i + 1

			if newLen != expectedLen {
				t.Errorf("invalid list lenght, got %d, expected %d", newLen, expectedLen)
			}

			tail := list.Tail()
			if tail != node {
				t.Errorf("invalid tail, got %v, expected %v", tail, node)
			}
		}
	})

	t.Run("iterate from tail", func(t *testing.T) {
		item := list.Tail()
		for i := len(testNodes) - 1; i >= 0; i-- {
			expectedItem := testNodes[i]

			if item != expectedItem {
				t.Errorf("invalid item, got %v, expedted, %v", item, expectedItem)
			}

			item = item.Next()
		}

		// last value should be nil to stop iteration
		if item != nil {
			t.Errorf("invalid item, got %v, expedted, %v", item, nil)
		}
	})

	t.Run("pop values", func(t *testing.T) {
		for i := len(testNodes) - 1; i >= 0; i-- {
			expedtedNode := testNodes[i]
			expectedLen := i

			node := list.PopTail()
			newLen := list.Len()

			if newLen != expectedLen {
				t.Errorf("invalid list lenght, got %d, expected %d", newLen, expectedLen)
			}

			if node != expedtedNode {
				t.Errorf("invalid head, got %v, expected %v", node, expedtedNode)
			}
		}

		// we "pop" all items, expedcted empty list
		if list.Len() != 0 {
			t.Errorf("invalid len, got %d, expected %d", list.Len(), 0)
		}
		if list.Head() != nil {
			t.Errorf("invalid head, got %v, expected %v", list.Head(), nil)
		}
		if list.Tail() != nil {
			t.Errorf("invalid tail, got %v, expected %v", list.Tail(), nil)
		}
	})
}

func TestMiddle(t *testing.T) {
	list := DoubleLinkedList{}
	for _, node := range testNodes {
		list.PushHead(node)
	}

	t.Run("insert values before", func(t *testing.T) {
		for _, item := range testNodes {
			newNode := &DoubleListNode{Value: 0}

			err := list.InsertBefore(item, newNode)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if item.Prev() != newNode {
				t.Errorf("invalid prev item, got %v, expected %v", item.Prev(), newNode)
			}
		}
	})

	t.Run("insert values after", func(t *testing.T) {
		for _, item := range testNodes {
			newNode := &DoubleListNode{Value: 0}

			err := list.InsertAfter(item, newNode)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if item.Next() != newNode {
				t.Errorf("invalid prev item, got %v, expected %v", item.Next(), newNode)
			}
		}
	})

	t.Run("remove values", func(t *testing.T) {
		err := list.Remove(testNodeToRemove)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		// iterate over list to check if node actually removed
		for i := list.Tail(); i != nil; i = i.Next() {
			if i == testNodeToRemove {
				t.Errorf("node has not been removed (iter from tail)")
			}
		}
		for i := list.Head(); i != nil; i = i.Prev() {
			if i == testNodeToRemove {
				t.Errorf("node has not been removed (iter from tail)")
			}
		}
	})
}

func TestMiddleWithNotExistingNode(t *testing.T) {
	list := DoubleLinkedList{}
	for _, node := range testNodes {
		list.PushHead(node)
	}

	t.Run("insert values before", func(t *testing.T) {
		newNode := &DoubleListNode{Value: 0}

		err := list.InsertBefore(testOtherNode, newNode)
		if err != ErrNotExists {
			t.Errorf("unexpected error, got %v, expected %v", err, ErrNotExists)
		}
	})

	t.Run("insert values after", func(t *testing.T) {
		newNode := &DoubleListNode{Value: 0}

		err := list.InsertAfter(testOtherNode, newNode)
		if err != ErrNotExists {
			t.Errorf("unexpected error, got %v, expected %v", err, ErrNotExists)
		}
	})

	t.Run("remove values", func(t *testing.T) {
		err := list.Remove(testOtherNode)
		if err != ErrNotExists {
			t.Errorf("unexpected error, got %v, expected %v", err, ErrNotExists)
		}
	})
}
