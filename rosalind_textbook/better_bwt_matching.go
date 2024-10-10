package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type BWTIndex struct {
	bwt             string
	firstOccurrence map[byte]int
	count           map[byte][]int
}

func newBWTIndex(bwt string) *BWTIndex {
	index := &BWTIndex{
		bwt:             bwt,
		firstOccurrence: make(map[byte]int),
		count:           make(map[byte][]int),
	}
	index.computeFirstOccurrence()
	index.computeCount()
	return index
}

func (index *BWTIndex) computeFirstOccurrence() {
	sorted := []byte(index.bwt)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	for i, char := range sorted {
		if _, exists := index.firstOccurrence[char]; !exists {
			index.firstOccurrence[char] = i
		}
	}
}

func (index *BWTIndex) computeCount() {
	for char := range index.firstOccurrence {
		index.count[char] = make([]int, len(index.bwt)+1)
	}
	for i, char := range index.bwt {
		for c := range index.count {
			index.count[c][i+1] = index.count[c][i]
		}
		index.count[byte(char)][i+1]++
	}
}

func (index *BWTIndex) betterBWMatching(pattern string) int {
	top := 0
	bottom := len(index.bwt) - 1

	for i := len(pattern) - 1; i >= 0; i-- {
		symbol := byte(pattern[i])
		if firstOcc, exists := index.firstOccurrence[symbol]; exists {
			top = firstOcc + index.count[symbol][top]
			bottom = firstOcc + index.count[symbol][bottom+1] - 1
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
	content, err := ioutil.ReadFile("rosalind_ba9m.txt")
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

	// Perform BetterBWMatching for each pattern
	results := make([]int, len(patterns))
	for i, pattern := range patterns {
		results[i] = index.betterBWMatching(pattern)
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
