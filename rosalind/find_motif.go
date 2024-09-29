package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function to find all starting positions of t as a substring of s
func findSubstringPositions(s, t string) []int {
	var positions []int
	tLen := len(t)
	for i := 0; i <= len(s)-tLen; i++ {
		if s[i:i+tLen] == t {
			positions = append(positions, i+1) // Convert to 1-based index
		}
	}
	return positions
}

func main() {
	// Open the file "rosalind_subs.txt"
	file, err := os.Open("rosalind_subs.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	var s, t string
	if scanner.Scan() {
		s = scanner.Text()
	}
	if scanner.Scan() {
		t = scanner.Text()
	}

	// Find all starting positions of t in s
	positions := findSubstringPositions(s, t)

	// Print the result
	for _, pos := range positions {
		fmt.Print(pos, " ")
	}
	fmt.Println()
}
