package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// patternToNumber converts a DNA string to a numerical representation
func patternToNumber(pattern string) int {
	nucleotideMap := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	result := 0
	for i := 0; i < len(pattern); i++ {
		result = 4*result + nucleotideMap[pattern[i]]
	}
	return result
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba1l.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read input from file
	scanner := bufio.NewScanner(file)

	// Read Pattern
	scanner.Scan()
	pattern := strings.TrimSpace(scanner.Text())

	// Convert pattern to number
	number := patternToNumber(pattern)

	// Print the result
	fmt.Println(number)
}
