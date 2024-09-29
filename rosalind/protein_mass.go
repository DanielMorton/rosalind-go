package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Monoisotopic mass table for amino acids
var massTable = map[byte]float64{
	'A': 71.03711, 'C': 103.00919, 'D': 115.02694, 'E': 129.04259,
	'F': 147.06841, 'G': 57.02146, 'H': 137.05891, 'I': 113.08406,
	'K': 128.09496, 'L': 113.08406, 'M': 131.04049, 'N': 114.04293,
	'P': 97.05276, 'Q': 128.05858, 'R': 156.10111, 'S': 87.03203,
	'T': 101.04768, 'V': 99.06841, 'W': 186.07931, 'Y': 163.06333,
}

// Function to calculate the total weight of the protein string
func calculateWeight(protein string) float64 {
	totalWeight := 0.0

	// Sum the weight of each amino acid in the protein string
	for i := 0; i < len(protein); i++ {
		totalWeight += massTable[protein[i]]
	}

	return totalWeight
}

func main() {
	// Read input from the file
	data, err := ioutil.ReadFile("rosalind_prtm.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Trim any surrounding whitespace and parse the protein string
	protein := strings.TrimSpace(string(data))

	// Calculate and output the total weight of the protein string
	totalWeight := calculateWeight(protein)
	fmt.Printf("%.3f\n", totalWeight)
}
