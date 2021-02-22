package structures

import (
	"testing"
)

var testStackValues = []int{1, 2, 3}

func TestStack(t *testing.T) {
	stack := NewStack()

	for _, item := range testStackValues {
		stack.Push(item)
	}

	if stack.Len() != len(testStackValues) {
		t.Errorf("invalid len, got %d, expected %d", stack.Len(), len(testStackValues))
	}

	// stack returns values in reverse order
	for i := len(testStackValues) - 1; i >= 0; i-- {
		expectedItem := testStackValues[i]

		item, err := stack.Pop()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if item != expectedItem {
			t.Errorf("invalid item, got %v, expected %v", item, expectedItem)
		}
	}

	// stach should be empty after we "pop" all values
	_, err := stack.Pop()
	if err != ErrEmptyStack {
		t.Errorf("unexpected error, got %v, expected %v", err, ErrEmptyStack)
	}
}
