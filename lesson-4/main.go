package main

import (
	"fmt"
	"gb-go-architecture/lesson-4/binary_search"
	"log"
	"os"
)

type TestCaseInt struct {
	name        string
	data        []int
	searchElem  int
	expectedIdx int
	expectedErr error
}

var (
	min  int = 0
	max  int = 10_000
	step int = 1

	resultFileUniq       string      = fmt.Sprintf("complexity_%d_%d_%d.csv", min, max, step)
	resultFileDuplicates string      = fmt.Sprintf("complexity_duplicates_%d_%d_%d.csv", min, max, step)
	resultFileMode       int         = os.O_WRONLY | os.O_CREATE
	resultFilePerm       os.FileMode = 0644
)

func initTestSlice(sliceLen int) (elems []int, lastElem int) {
	elems = make([]int, sliceLen)
	for i := 0; i < sliceLen; i++ {
		lastElem = i + 1
	}
	if len(elems) > 0 {
		lastElem = elems[len(elems)-1]
	}
	return elems, lastElem
}

func testComplexityUnuqElems(resultFile *os.File) {
	testSlice, lastElem := initTestSlice(min)
	for i := min; i < max; i += step {
		// increase tested slice
		initialLen := len(testSlice)
		targetLen := len(testSlice) + step
		for i := initialLen; i < targetLen; i++ {
			lastElem = i + 1
			testSlice = append(testSlice, lastElem)
		}

		// test search functions
		_, steps1, _ := binary_search.SearchInt(testSlice, lastElem)
		_, steps2, _ := binary_search.SearchFirstInt(testSlice, lastElem)

		// report results
		dataStr := fmt.Sprintf("%d,%v,%v\n", len(testSlice), steps1, steps2)
		_, err := resultFile.WriteString(dataStr)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func testComplexityDuplicateElems(resultFile *os.File) {
	const duplicatesPartDivider = 5

	testSlice, lastElem := initTestSlice(min)
	for i := min; i < max; i += step {
		// increase tested slice
		targetLen := len(testSlice) + step
		duplicatesNewLen := targetLen - targetLen/duplicatesPartDivider
		// 1) override previous duplicates
		duplicatesOldLen := len(testSlice) - len(testSlice)/duplicatesPartDivider
		testSlice = testSlice[:duplicatesOldLen]
		for i := duplicatesOldLen; i < duplicatesNewLen; i++ {
			lastElem = i + 1
			testSlice = append(testSlice, lastElem)
		}
		// 2) append new duplicates
		for i := duplicatesNewLen; i < targetLen; i++ {
			testSlice = append(testSlice, lastElem)
		}

		// test search functions
		_, steps1, _ := binary_search.SearchInt(testSlice, lastElem)
		_, steps2, _ := binary_search.SearchFirstInt(testSlice, lastElem)

		// report results
		dataStr := fmt.Sprintf("%d,%v,%v\n", len(testSlice), steps1, steps2)
		_, err := resultFile.WriteString(dataStr)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	// file for results
	resultFile1, err := os.OpenFile(resultFileUniq, resultFileMode, resultFilePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer resultFile1.Close()

	// slices contains unique elements
	testComplexityUnuqElems(resultFile1)

	// file for results
	resultFile2, err := os.OpenFile(resultFileDuplicates, resultFileMode, resultFilePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer resultFile2.Close()

	// slices contains duplicates (1/4 of len)
	testComplexityDuplicateElems(resultFile2)
}
