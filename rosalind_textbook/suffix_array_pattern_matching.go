package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type suffixArray struct {
	text string
	sa   []int
}

func newSuffixArray(text string) *suffixArray {
	sa := &suffixArray{
		text: text,
		sa:   make([]int, len(text)),
	}
	for i := range sa.sa {
		sa.sa[i] = i
	}
	sort.Slice(sa.sa, func(i, j int) bool {
		return sa.less(sa.sa[i], sa.sa[j])
	})
	return sa
}

func (sa *suffixArray) less(i, j int) bool {
	for k := 0; k < len(sa.text)-i && k < len(sa.text)-j; k++ {
		if sa.text[i+k] != sa.text[j+k] {
			return sa.text[i+k] < sa.text[j+k]
		}
	}
	return i > j
}

func (sa *suffixArray) binarySearch(pattern string) (int, int) {
	left, right := 0, len(sa.sa)-1
	start, end := -1, -1

	// Find start
	for left <= right {
		mid := (left + right) / 2
		suffix := sa.text[sa.sa[mid]:]
		if strings.HasPrefix(suffix, pattern) {
			start = mid
			right = mid - 1
		} else if suffix < pattern {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if start == -1 {
		return -1, -1
	}

	// Find end
	left, right = start, len(sa.sa)-1
	for left <= right {
		mid := (left + right) / 2
		suffix := sa.text[sa.sa[mid]:]
		if strings.HasPrefix(suffix, pattern) {
			end = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return start, end
}

func multiplePatternMatching(text string, patterns []string) []int {
	sa := newSuffixArray(text)
	result := make(map[int]bool)

	for _, pattern := range patterns {
		start, end := sa.binarySearch(pattern)
		if start != -1 {
			for i := start; i <= end; i++ {
				result[sa.sa[i]] = true
			}
		}
	}

	positions := make([]int, 0, len(result))
	for pos := range result {
		positions = append(positions, pos)
	}
	sort.Ints(positions)

	return positions
}

func main() {
	// Read input from file
	content, err := ioutil.ReadFile("rosalind_ba9h.txt")
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

	// Find all starting positions of patterns in text
	positions := multiplePatternMatching(text, patterns)

	// Print the result
	for i, pos := range positions {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(pos)
	}
	fmt.Println()
}
