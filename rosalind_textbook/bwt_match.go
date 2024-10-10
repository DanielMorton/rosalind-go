package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type BWTIndex struct {
	bwt             string
	firstColumn     string
	occurrences     map[byte][]int
	firstOccurrence map[byte]int
}

func newBWTIndex(bwt string) *BWTIndex {
	index := &BWTIndex{
		bwt:             bwt,
		firstColumn:     getSortedBWT(bwt),
		occurrences:     make(map[byte][]int),
		firstOccurrence: make(map[byte]int),
	}
	index.computeOccurrences()
	index.computeFirstOccurrence()
	return index
}

func getSortedBWT(bwt string) string {
	sorted := []byte(bwt)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	return string(sorted)
}

func (index *BWTIndex) computeOccurrences() {
	for i, char := range index.bwt {
		if _, exists := index.occurrences[byte(char)]; !exists {
			index.occurrences[byte(char)] = make([]int, len(index.bwt)+1)
		}
		for ch := range index.occurrences {
			if ch == byte(char) {
				index.occurrences[ch][i+1] = index.occurrences[ch][i] + 1
			} else {
				index.occurrences[ch][i+1] = index.occurrences[ch][i]
			}
		}
	}
}

func (index *BWTIndex) computeFirstOccurrence() {
	for i, char := range index.firstColumn {
		if _, exists := index.firstOccurrence[byte(char)]; !exists {
			index.firstOccurrence[byte(char)] = i
		}
	}
}

func (index *BWTIndex) bwMatching(pattern string) int {
	top := 0
	bottom := len(index.bwt) - 1

	for i := len(pattern) - 1; i >= 0; i-- {
		symbol := pattern[i]
		if occurrences, exists := index.occurrences[symbol]; exists {
			top = index.firstOccurrence[symbol] + occurrences[top]
			bottom = index.firstOccurrence[symbol] + occurrences[bottom+1] - 1
			if top > bottom {
				return 0
			}
		} else {
			return 0
		}
	}

	return bottom - top + 1
}

func main() {
	// Read input from file
	content, err := ioutil.ReadFile("rosalind_ba9l.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the content into BWT and patterns
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	if len(lines) != 2 {
		fmt.Println("Invalid input format")
		return
	}

	bwt := lines[0]
	patterns := strings.Split(lines[1], " ")

	// Create BWT index
	index := newBWTIndex(bwt)

	// Perform BWMatching for each pattern
	results := make([]int, len(patterns))
	for i, pattern := range patterns {
		results[i] = index.bwMatching(pattern)
	}

	// Print the results
	for i, count := range results {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(count)
	}
	fmt.Println()
}
