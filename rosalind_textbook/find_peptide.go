package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AminoAcid struct {
	Mass   int
	Symbol string
}

var aminoAcids = []AminoAcid{
	{71, "A"}, {103, "C"}, {115, "D"}, {129, "E"}, {147, "F"},
	{57, "G"}, {137, "H"}, {113, "I"}, {128, "K"}, {113, "L"},
	{131, "M"}, {114, "N"}, {97, "P"}, {128, "Q"}, {156, "R"},
	{87, "S"}, {101, "T"}, {99, "V"}, {186, "W"}, {163, "Y"},
	{4, "X"}, {5, "Z"}, // Imaginary amino acids X and Z
}

func main() {
	spectralVector, proteome := readInput("rosalind_ba11f.txt")
	result := findOptimalPeptide(spectralVector, proteome)
	fmt.Println(result)
}

func readInput(filename string) ([]int, string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read spectral vector
	scanner.Scan()
	line := strings.TrimSpace(scanner.Text())
	var spectralVector []int
	spectralVector = append(spectralVector, 0) // Add a 0 at index 0 to align with 1-based indexing
	for _, s := range strings.Fields(line) {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Invalid input: %s", s))
		}
		spectralVector = append(spectralVector, num)
	}

	// Read proteome
	scanner.Scan()
	proteome := strings.TrimSpace(scanner.Text())

	return spectralVector, proteome
}

func findOptimalPeptide(spectralVector []int, proteome string) string {
	targetMass := len(spectralVector) - 1 // Subtract 1 because spectralVector is 1-indexed
	bestScore := -1000000
	var bestPeptide string

	for i := 0; i < len(proteome); i++ {
		mass := 0
		for j := i; j < len(proteome); j++ {
			mass += findMass(string(proteome[j]))
			if mass == targetMass {
				peptide := proteome[i : j+1]
				score := scorePeptide(peptide, spectralVector)
				if score > bestScore {
					bestScore = score
					bestPeptide = peptide
				}
				break // No need to continue this inner loop
			} else if mass > targetMass {
				break // Stop if we've exceeded the target mass
			}
		}
	}

	return bestPeptide
}

func scorePeptide(peptide string, spectralVector []int) int {
	prefixMass := 0
	score := 0
	for _, aa := range peptide {
		prefixMass += findMass(string(aa))
		score += spectralVector[prefixMass]
	}
	return score
}

func findMass(symbol string) int {
	for _, aa := range aminoAcids {
		if aa.Symbol == symbol {
			return aa.Mass
		}
	}
	panic(fmt.Sprintf("Unknown amino acid: %s", symbol))
}
