package binary_search

import (
	"errors"
)

var (
	ErrNotFound = errors.New("Element not found")
)

func checkEdgeCases(elems []int, searchedElem int) (idx, steps int, err error) {
	// all operations:
	// get lenght, get elem by index, expression eveluation
	// are constant
	steps = 1

	// elems is empty
	if len(elems) == 0 {
		return idx, steps, ErrNotFound
	}
	first := 0
	last := len(elems) - 1
	// if elem > last or elem < first => elem does not exists in sorted list
	if searchedElem < elems[first] || searchedElem > elems[last] {
		return idx, steps, ErrNotFound
	}

	return 0, steps, nil
}

func SearchInt(elems []int, searchedElem int) (idx, steps int, err error) {
	idx, ecSteps, err := checkEdgeCases(elems, searchedElem)
	if err != nil {
		return idx, ecSteps, err
	}

	first := 0
	last := len(elems) - 1
	// regular case
	for offset := last - first; ; offset = last - first {
		// all operations inside 'for' has constant complexity
		// so, we just count iterations
		steps++
		middle := first + offset/2
		if searchedElem == elems[middle] {
			return middle, steps, nil
		}
		if searchedElem > elems[middle] {
			first = middle + 1
		} else {
			last = middle
		}

		// we cannot split slice anymore
		if offset < 1 {
			break
		}
	}

	return idx, steps, ErrNotFound
}

func SearchFirstInt(elems []int, searchedElem int) (idx, steps int, err error) {
	idx, ecSteps, err := checkEdgeCases(elems, searchedElem)
	if err != nil {
		return idx, ecSteps, err
	}

	first := 0
	last := len(elems) - 1
	var middle int
	// regular case
	for offset := last - first; ; offset = last - first {
		// all operations inside 'for' has constant complexity
		// so, we just count iterations
		steps++
		middle = first + offset/2
		if searchedElem > elems[middle] {
			first = middle + 1
		} else {
			last = middle
		}

		// we cannot split slice anymore
		if offset < 1 {
			break
		}
	}

	if searchedElem == elems[middle] {
		return middle, steps, nil
	}

	return idx, steps, ErrNotFound
}
