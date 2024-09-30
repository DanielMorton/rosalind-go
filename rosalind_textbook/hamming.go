package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// hammingDistance calculates the Hamming distance between two DNA strings
func hammingDistance(s1, s2 string) (int, error) {
	if len(s1) != len(s2) {
		return 0, fmt.Errorf("strings must be of equal length")
	}

	distance := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			distance++
		}
	}

	return distance, nil
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba1g.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read input from file
	scanner := bufio.NewScanner(file)

	// Read first DNA string
	scanner.Scan()
	dna1 := strings.TrimSpace(scanner.Text())

	// Read second DNA string
	scanner.Scan()
	dna2 := strings.TrimSpace(scanner.Text())

	// Calculate Hamming distance
	distance, err := hammingDistance(dna1, dna2)
	if err != nil {
		fmt.Println("Error calculating Hamming distance:", err)
		return
	}

	// Print the result
	fmt.Println(distance)
}
