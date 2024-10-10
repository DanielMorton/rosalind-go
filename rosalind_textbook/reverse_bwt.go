package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func inverseBWT(transform string) string {
	n := len(transform)

	// Create a sorted version of the transform
	sorted := make([]byte, n)
	copy(sorted, transform)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	// Create first-to-last mapping
	ftl := make([]int, n)
	counts := make(map[byte]int)
	for i := 0; i < n; i++ {
		char := transform[i]
		ftl[i] = counts[char]
		counts[char]++
	}

	// Reconstruct the original string
	result := make([]byte, n)
	idx := 0 // Start with the '$' character
	for i := n - 1; i >= 0; i-- {
		result[i] = transform[idx]
		idx = findNthOccurrence(sorted, transform[idx], ftl[idx])
	}

	return string(result)
}

func findNthOccurrence(s []byte, char byte, n int) int {
	count := 0
	for i, c := range s {
		if c == char {
			if count == n {
				return i
			}
			count++
		}
	}
	return -1 // This should never happen if the input is valid
}

func main() {
	// Read input from file
	content, err := ioutil.ReadFile("rosalind_ba9j.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Trim whitespace and newlines
	transform := strings.TrimSpace(string(content))

	// Reconstruct the original string
	original := inverseBWT(transform)

	// Print the result in both required formats
	fmt.Println(original[1:] + string(original[0]))
}
