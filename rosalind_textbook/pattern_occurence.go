package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// findPatternOccurrences returns all starting positions of Pattern in Genome
func findPatternOccurrences(pattern, genome string) []int {
	var positions []int
	for i := 0; i <= len(genome)-len(pattern); i++ {
		if genome[i:i+len(pattern)] == pattern {
			positions = append(positions, i)
		}
	}
	return positions
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba1d.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read input from file
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	pattern := strings.TrimSpace(scanner.Text())
	scanner.Scan()
	genome := strings.TrimSpace(scanner.Text())

	// Find all occurrences of the pattern
	positions := findPatternOccurrences(pattern, genome)

	// Print the result
	for i, pos := range positions {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(pos)
	}
	fmt.Println()
}
