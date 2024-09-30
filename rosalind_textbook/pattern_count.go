package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func patternCount(text, pattern string) int {
	count := 0
	for i := 0; i <= len(text)-len(pattern); i++ {
		if text[i:i+len(pattern)] == pattern {
			count++
		}
	}
	return count
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba1a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read input from file
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := strings.TrimSpace(scanner.Text())
	scanner.Scan()
	pattern := strings.TrimSpace(scanner.Text())

	// Compute the count
	result := patternCount(text, pattern)

	// Print the result
	fmt.Println(result)
}
