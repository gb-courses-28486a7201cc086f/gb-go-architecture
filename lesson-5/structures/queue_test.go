package structures

import (
	"testing"
)

var testQueueValues = []int{1, 2, 3}

func TestQueue(t *testing.T) {
	queue := NewQueue()

	for _, item := range testQueueValues {
		queue.Push(item)
	}

	if queue.Len() != len(testQueueValues) {
		t.Errorf("invalid len, got %d, expected %d", queue.Len(), len(testQueueValues))
	}

	for _, expectedItem := range testQueueValues {
		item, err := queue.Pop()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if item != expectedItem {
			t.Errorf("invalid item, got %v, expected %v", item, expectedItem)
		}
	}

	// stach should be empty after we "pop" all values
	_, err := queue.Pop()
	if err != ErrEmptyQueue {
		t.Errorf("unexpected error, got %v, expected %v", err, ErrEmptyQueue)
	}
}
