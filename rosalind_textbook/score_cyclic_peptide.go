package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var aminoAcidMass = map[string]int{
	"G": 57, "A": 71, "S": 87, "P": 97, "V": 99,
	"T": 101, "C": 103, "I": 113, "L": 113, "N": 114,
	"D": 115, "K": 128, "Q": 128, "E": 129, "M": 131,
	"H": 137, "F": 147, "R": 156, "Y": 163, "W": 186,
}

func mass(peptide string) int {
	total := 0
	for _, aa := range peptide {
		total += aminoAcidMass[string(aa)]
	}
	return total
}

func cyclospectrum(peptide string) []int {
	n := len(peptide)
	extendedPeptide := peptide + peptide[:n-1]
	spectrum := []int{0, mass(peptide)} // Include 0 and total mass

	for length := 1; length < n; length++ {
		for i := 0; i < n; i++ {
			subPeptide := extendedPeptide[i : i+length]
			spectrum = append(spectrum, mass(subPeptide))
		}
	}

	return spectrum
}

func score(peptide string, spectrum []int) int {
	theoreticalSpectrum := cyclospectrum(peptide)

	// Count frequencies in both spectra
	theoreticalFreq := make(map[int]int)
	experimentalFreq := make(map[int]int)

	for _, mass := range theoreticalSpectrum {
		theoreticalFreq[mass]++
	}
	for _, mass := range spectrum {
		experimentalFreq[mass]++
	}

	// Calculate score
	score := 0
	for mass, count := range theoreticalFreq {
		if experimentalCount, exists := experimentalFreq[mass]; exists {
			score += min(count, experimentalCount)
		}
	}

	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	file, err := os.Open("rosalind_ba4f.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read peptide
	scanner.Scan()
	peptide := scanner.Text()

	// Read spectrum
	scanner.Scan()
	spectrumStr := strings.Fields(scanner.Text())
	spectrum := make([]int, len(spectrumStr))
	for i, s := range spectrumStr {
		spectrum[i], _ = strconv.Atoi(s)
	}

	result := score(peptide, spectrum)
	fmt.Println(result)
}
