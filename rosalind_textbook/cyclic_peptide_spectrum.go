package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var aminoAcidMass = map[string]int{
	"G": 57, "A": 71, "S": 87, "P": 97, "V": 99,
	"T": 101, "C": 103, "I": 113, "L": 113, "N": 114,
	"D": 115, "K": 128, "Q": 128, "E": 129, "M": 131,
	"H": 137, "F": 147, "R": 156, "Y": 163, "W": 186,
}

func calculatePeptideMass(peptide string) int {
	mass := 0
	for _, aa := range peptide {
		mass += aminoAcidMass[string(aa)]
	}
	return mass
}

func generateCyclicSubpeptides(peptide string) []string {
	var subpeptides []string
	n := len(peptide)
	extendedPeptide := peptide + peptide[:n-1]

	for length := 1; length < n; length++ {
		for start := 0; start < n; start++ {
			subpeptide := extendedPeptide[start : start+length]
			subpeptides = append(subpeptides, subpeptide)
		}
	}

	// Add the full peptide only once
	subpeptides = append(subpeptides, peptide)

	return subpeptides
}

func generateTheoreticalSpectrum(peptide string) []int {
	var spectrum []int
	spectrum = append(spectrum, 0) // Mass of empty peptide

	subpeptides := generateCyclicSubpeptides(peptide)
	for _, subpeptide := range subpeptides {
		spectrum = append(spectrum, calculatePeptideMass(subpeptide))
	}

	sort.Ints(spectrum)
	return spectrum
}

func main() {
	file, err := os.Open("rosalind_ba4c.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	peptide := strings.TrimSpace(scanner.Text())

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	spectrum := generateTheoreticalSpectrum(peptide)

	// Print the spectrum
	for i, mass := range spectrum {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(mass)
	}
	fmt.Println()
}
