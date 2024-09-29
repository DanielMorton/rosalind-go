package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file named "rosalind_revc.txt"
	file, err := os.Open("rosalind_revc.txt")
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

	// Reverse the DNA string and find the complement
	reverseComplement := reverseAndComplement(dna)

	// Output the reverse complement string
	fmt.Println(reverseComplement)
}

// Function to get the reverse complement of a DNA string
func reverseAndComplement(dna string) string {
	// Create a map for complementing DNA bases
	complements := map[byte]byte{
		'A': 'T',
		'T': 'A',
		'C': 'G',
		'G': 'C',
	}

	// Create a slice to store the reverse complement
	reverseComplement := make([]byte, len(dna))

	// Iterate over the DNA string in reverse order
	for i := 0; i < len(dna); i++ {
		// Get the complement of each nucleotide
		reverseComplement[i] = complements[dna[len(dna)-1-i]]
	}

	return string(reverseComplement)
}
