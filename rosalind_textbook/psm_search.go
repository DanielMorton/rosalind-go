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

type PSM struct {
	Peptide string
	Vector  []int
}

func main() {
	spectralVectors, proteome, threshold := readInput("rosalind_ba11g.txt")
	result := psmSearch(spectralVectors, proteome, threshold)
	for _, psm := range result {
		fmt.Println(psm.Peptide)
	}
}

func readInput(filename string) ([][]int, string, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var spectralVectors [][]int
	var proteome string
	var threshold int

	// Read spectral vectors
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		// Check if this line is the proteome sequence
		if _, err := strconv.Atoi(strings.Fields(line)[0]); err != nil {
			proteome = line
			break
		}
		// Otherwise, it's a spectral vector
		var vector []int
		vector = append(vector, 0) // Add a 0 at index 0 to align with 1-based indexing
		for _, s := range strings.Fields(line) {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(fmt.Sprintf("Invalid input in spectral vector: %s", s))
			}
			vector = append(vector, num)
		}
		spectralVectors = append(spectralVectors, vector)
	}

	// Read threshold
	scanner.Scan()
	threshold, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		panic("Invalid threshold value")
	}

	return spectralVectors, proteome, threshold
}

func psmSearch(spectralVectors [][]int, proteome string, threshold int) []PSM {
	var psmSet []PSM
	for _, vector := range spectralVectors {
		peptide := peptideIdentification(vector, proteome)
		score := scorePeptide(peptide, vector)
		if score >= threshold {
			psmSet = append(psmSet, PSM{Peptide: peptide, Vector: vector})
		}
	}
	return uniquePSMs(psmSet)
}

func peptideIdentification(spectralVector []int, proteome string) string {
	targetMass := len(spectralVector) - 1
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
				break
			} else if mass > targetMass {
				break
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

func uniquePSMs(psms []PSM) []PSM {
	uniqueMap := make(map[string]bool)
	var uniquePSMs []PSM
	for _, psm := range psms {
		if !uniqueMap[psm.Peptide] {
			uniqueMap[psm.Peptide] = true
			uniquePSMs = append(uniquePSMs, psm)
		}
	}
	return uniquePSMs
}
