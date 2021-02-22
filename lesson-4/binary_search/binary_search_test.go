package binary_search

import "testing"

type TestCaseInt struct {
	name        string
	data        []int
	searchElem  int
	expectedIdx int
	expectedErr error
}

var (
	testDataInt []*TestCaseInt = []*TestCaseInt{
		{"empty list", []int{}, 100, 0, ErrNotFound},
		{"too small elem", []int{4, 7, 9}, 1, 0, ErrNotFound},
		{"too big elem", []int{4, 7, 9}, 10, 0, ErrNotFound},
		{"single existing elem", []int{100}, 100, 0, nil},
		{"single not existing elem", []int{1}, 100, 0, ErrNotFound},
		{"odd len list", []int{4, 7, 9, 34, 56}, 34, 3, nil},
		{"first elem in odd len list", []int{4, 7, 9, 34, 56}, 4, 0, nil},
		{"last elem in odd len list", []int{4, 7, 9, 34, 56}, 56, 4, nil},
		{"not existing elem in odd len list", []int{4, 7, 9, 34, 56}, 50, 0, ErrNotFound},
		{"even len list", []int{4, 7, 9, 34, 56, 87}, 34, 3, nil},
		{"first elem in even len list", []int{4, 7, 9, 34, 56, 87}, 4, 0, nil},
		{"last elem in even len list", []int{4, 7, 9, 34, 56, 87}, 87, 5, nil},
		{"not existing elem in even len list", []int{4, 7, 9, 34, 56, 87}, 50, 0, ErrNotFound},
	}

	testDataDuplicatesInt []*TestCaseInt = []*TestCaseInt{
		{"duplicated single value odd len", []int{100, 100, 100}, 100, 1, nil},
		{"duplicated single value even len", []int{100, 100, 100, 100}, 100, 1, nil},
		{"duplicated elem left side odd len list", []int{4, 7, 9, 9, 9, 34, 34, 34, 56}, 9, 4, nil},
		{"duplicated elem right side odd len list", []int{4, 7, 9, 9, 9, 34, 34, 34, 56}, 9, 4, nil},
		{"duplicated elem left side even len list", []int{4, 7, 9, 9, 9, 34, 34, 34, 56, 87}, 34, 7, nil},
		{"duplicated elem right side even len list", []int{4, 7, 9, 9, 9, 34, 34, 34, 56, 87}, 34, 7, nil},
	}

	testDataDuplicatesFirstInt []*TestCaseInt = []*TestCaseInt{
		{"duplicated single value odd len", []int{100, 100, 100}, 100, 0, nil},
		{"duplicated single value even len", []int{100, 100, 100, 100}, 100, 0, nil},
		{"duplicated elem left side odd len list", []int{4, 7, 9, 9, 9, 34, 34, 34, 56}, 9, 2, nil},
		{"duplicated elem right side odd len list", []int{4, 7, 9, 9, 9, 34, 34, 34, 56}, 9, 2, nil},
		{"duplicated elem left side even len list", []int{4, 7, 9, 9, 9, 34, 34, 34, 56, 87}, 34, 5, nil},
		{"duplicated elem right side even len list", []int{4, 7, 9, 9, 9, 34, 34, 34, 56, 87}, 34, 5, nil},
	}
)

func TestSearchInt(t *testing.T) {
	for _, testCase := range testDataInt {
		t.Run(testCase.name, func(t *testing.T) {
			idx, _, err := SearchInt(testCase.data, testCase.searchElem)
			if idx != testCase.expectedIdx {
				t.Errorf("Invalid index: got %d, expected %d", idx, testCase.expectedIdx)
			}
			if err != testCase.expectedErr {
				t.Errorf("Invalid err: got %v, expected %v", err, testCase.expectedErr)
			}
		})
	}

	// test search duplicated elements
	// SearchInt can return first match, even it is not minimal index
	for _, testCase := range testDataDuplicatesInt {
		t.Run(testCase.name, func(t *testing.T) {
			idx, _, err := SearchInt(testCase.data, testCase.searchElem)
			if idx != testCase.expectedIdx {
				t.Errorf("Invalid index: got %d, expected %d", idx, testCase.expectedIdx)
			}
			if err != testCase.expectedErr {
				t.Errorf("Invalid err: got %v, expected %v", err, testCase.expectedErr)
			}
		})
	}
}
func TestSearchFirstInt(t *testing.T) {
	for _, testCase := range testDataInt {
		t.Run(testCase.name, func(t *testing.T) {
			idx, _, err := SearchFirstInt(testCase.data, testCase.searchElem)
			if idx != testCase.expectedIdx {
				t.Errorf("Invalid index: got %d, expected %d", idx, testCase.expectedIdx)
			}
			if err != testCase.expectedErr {
				t.Errorf("Invalid err: got %v, expected %v", err, testCase.expectedErr)
			}
		})
	}

	// test search duplicated elements
	// SearchFirstInt should return minimal index
	for _, testCase := range testDataDuplicatesFirstInt {
		t.Run(testCase.name, func(t *testing.T) {
			idx, _, err := SearchFirstInt(testCase.data, testCase.searchElem)
			if idx != testCase.expectedIdx {
				t.Errorf("Invalid index: got %d, expected %d", idx, testCase.expectedIdx)
			}
			if err != testCase.expectedErr {
				t.Errorf("Invalid err: got %v, expected %v", err, testCase.expectedErr)
			}
		})
	}
}
