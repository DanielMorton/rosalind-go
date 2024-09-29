package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Function to calculate the number of possible RNA strings
func countRNAPossibilities(protein string) int {
	// Codon possibilities for each amino acid
	codonTable := map[byte]int{
		'A': 4, 'C': 2, 'D': 2, 'E': 2, 'F': 2,
		'G': 4, 'H': 2, 'I': 3, 'K': 2, 'L': 6,
		'M': 1, 'N': 2, 'P': 4, 'Q': 2, 'R': 6,
		'S': 6, 'T': 4, 'V': 4, 'W': 1, 'Y': 2,
	}

	mod := 1000000
	result := 1

	// Multiply the possibilities for each amino acid in the protein
	for i := 0; i < len(protein); i++ {
		result = (result * codonTable[protein[i]]) % mod
	}

	// Multiply by 3 for the stop codon
	result = (result * 3) % mod

	return result
}

func main() {
	// Read the input from the file
	data, err := ioutil.ReadFile("rosalind_mrna.txt")
	if err != nil {
		log.Fatal(err)
	}

	// The protein string from the file
	protein := strings.TrimSpace(string(data))

	// Output the result
	fmt.Println(countRNAPossibilities(protein))
}
