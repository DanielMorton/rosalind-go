package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// complementBase returns the complement of a given nucleotide
func complementBase(base byte) byte {
	switch base {
	case 'A':
		return 'T'
	case 'T':
		return 'A'
	case 'C':
		return 'G'
	case 'G':
		return 'C'
	default:
		return base // Return the same base if it's not A, T, C, or G
	}
}

// reverseComplement returns the reverse complement of a DNA string
func reverseComplement(pattern string) string {
	complement := make([]byte, len(pattern))
	for i := 0; i < len(pattern); i++ {
		complement[len(pattern)-1-i] = complementBase(pattern[i])
	}
	return string(complement)
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba1c.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read input from file
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	pattern := strings.TrimSpace(scanner.Text())

	// Find the reverse complement
	result := reverseComplement(pattern)

	// Print the result
	fmt.Println(result)
}
