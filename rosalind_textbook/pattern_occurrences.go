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
	occ             map[byte][]int
	suffixArray     []int
}

func newBWTIndex(text string) *BWTIndex {
	bwt, suffixArray := constructBWTAndSuffixArray(text)
	index := &BWTIndex{
		bwt:             bwt,
		firstOccurrence: make(map[byte]int),
		occ:             make(map[byte][]int),
		suffixArray:     suffixArray,
	}
	index.computeFirstOccurrence()
	index.computeOcc()
	return index
}

func constructBWTAndSuffixArray(text string) (string, []int) {
	text += "$"
	n := len(text)
	sa := make([]int, n)
	for i := range sa {
		sa[i] = i
	}
	sort.Slice(sa, func(i, j int) bool {
		for k := 0; k < n; k++ {
			a := (sa[i] + k) % n
			b := (sa[j] + k) % n
			if text[a] != text[b] {
				return text[a] < text[b]
			}
		}
		return false
	})
	bwt := make([]byte, n)
	for i, j := range sa {
		if j == 0 {
			bwt[i] = text[n-1]
		} else {
			bwt[i] = text[j-1]
		}
	}
	return string(bwt), sa
}

func (index *BWTIndex) computeFirstOccurrence() {
	for i, char := range index.bwt {
		if _, exists := index.firstOccurrence[byte(char)]; !exists {
			index.firstOccurrence[byte(char)] = i
		}
	}
}

func (index *BWTIndex) computeOcc() {
	for char := range index.firstOccurrence {
		index.occ[char] = make([]int, len(index.bwt)+1)
	}
	for i, char := range index.bwt {
		for c := range index.occ {
			index.occ[c][i+1] = index.occ[c][i]
		}
		index.occ[byte(char)][i+1]++
	}
}

func (index *BWTIndex) bwMatching(pattern string) []int {
	top := 0
	bottom := len(index.bwt) - 1
	for i := len(pattern) - 1; i >= 0; i-- {
		symbol := byte(pattern[i])
		if firstOcc, exists := index.firstOccurrence[symbol]; exists {
			top = firstOcc + index.occ[symbol][top]
			bottom = firstOcc + index.occ[symbol][bottom+1] - 1
			if top > bottom {
				return nil
			}
		} else {
			return nil
		}
	}
	return index.locateOccurrences(top, bottom)
}

func (index *BWTIndex) locateOccurrences(top, bottom int) []int {
	occurrences := make([]int, 0, bottom-top+1)
	for i := top; i <= bottom; i++ {
		occurrences = append(occurrences, index.suffixArray[i])
	}
	return occurrences
}

func multiplePatternMatching(text string, patterns []string) []int {
	index := newBWTIndex(text)
	allOccurrences := make(map[int]bool)
	for _, pattern := range patterns {
		occurrences := index.bwMatching(pattern)
		for _, occ := range occurrences {
			allOccurrences[occ] = true
		}
	}
	result := make([]int, 0, len(allOccurrences))
	for occ := range allOccurrences {
		result = append(result, occ)
	}
	sort.Ints(result)
	return result
}

func main() {
	// Read input from file
	content, err := ioutil.ReadFile("rosalind_ba9n.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the content into text and patterns
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	if len(lines) < 2 {
		fmt.Println("Invalid input format")
		return
	}

	text := lines[0]
	patterns := lines[1:]

	// Find all occurrences of patterns in text
	occurrences := multiplePatternMatching(text, patterns)

	// Print the results
	for i, occ := range occurrences {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(occ)
	}
	fmt.Println()
}
