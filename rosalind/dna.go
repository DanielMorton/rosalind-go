package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("rosalind_dna.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	var dna string
	for scanner.Scan() {
		dna += scanner.Text()
	}

	// Check for any error during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Initialize counters
	var aCount, cCount, gCount, tCount int

	// Count occurrences of 'A', 'C', 'G', 'T'
	for _, nucleotide := range dna {
		switch nucleotide {
		case 'A':
			aCount++
		case 'C':
			cCount++
		case 'G':
			gCount++
		case 'T':
			tCount++
		}
	}

	// Output the counts
	fmt.Printf("%d %d %d %d\n", aCount, cCount, gCount, tCount)
}
