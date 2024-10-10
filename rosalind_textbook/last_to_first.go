package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func lastToFirst(transform string, i int) int {
	n := len(transform)

	// Create a sorted version of the transform (FirstColumn)
	firstColumn := make([]byte, n)
	copy(firstColumn, transform)
	sort.Slice(firstColumn, func(i, j int) bool {
		return firstColumn[i] < firstColumn[j]
	})

	// Count occurrences of each character before position i in LastColumn
	counts := make(map[byte]int)
	for j := 0; j < i; j++ {
		counts[transform[j]]++
	}

	// Find the character at position i in LastColumn
	char := transform[i]

	// Find its position in FirstColumn
	return findNthOccurrence(firstColumn, char, counts[char])
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
	content, err := ioutil.ReadFile("rosalind_ba9k.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the content into transform and i
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	if len(lines) != 2 {
		fmt.Println("Invalid input format")
		return
	}

	transform := lines[0]
	i, err := strconv.Atoi(strings.TrimSpace(lines[1]))
	if err != nil {
		fmt.Println("Invalid integer input:", err)
		return
	}

	// Calculate the Last-to-First mapping
	result := lastToFirst(transform, i)

	// Print the result
	fmt.Println(result)
}
