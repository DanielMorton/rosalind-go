package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Amino acid mass table
var aminoAcidMass = map[string]float64{
	"A": 71.03711, "C": 103.00919, "D": 115.02694, "E": 129.04259,
	"F": 147.06841, "G": 57.02146, "H": 137.05891, "I": 113.08406,
	"K": 128.09496, "L": 113.08406, "M": 131.04049, "N": 114.04293,
	"P": 97.05276, "Q": 128.05858, "R": 156.10111, "S": 87.03203,
	"T": 101.04768, "V": 99.06841, "W": 186.07931, "Y": 163.06333,
}

// Find the closest amino acid mass to the given weight difference
func findClosestMass(weightDiff float64, tolerance float64) (string, bool) {
	for aa, mass := range aminoAcidMass {
		if math.Abs(weightDiff-mass) < tolerance {
			return aa, true
		}
	}
	return "", false
}

// reverseEngineerProtein reconstructs the protein string from the prefix spectrum
func reverseEngineerProtein(prefixSpectrum []float64) string {
	n := len(prefixSpectrum)
	if n == 0 {
		return ""
	}

	var protein strings.Builder
	tolerance := 0.001

	for i := 1; i < n; i++ {
		weightDiff := prefixSpectrum[i] - prefixSpectrum[i-1]
		aa, found := findClosestMass(weightDiff, tolerance)
		if !found {
			fmt.Println("Error: Cannot match the weight difference.")
			return ""
		}
		protein.WriteString(aa)
	}
	return protein.String()
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_spec.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var prefixSpectrum []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		weight, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error parsing weight:", err)
			return
		}
		prefixSpectrum = append(prefixSpectrum, weight)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Reconstruct the protein string
	protein := reverseEngineerProtein(prefixSpectrum)

	// Print the result
	fmt.Println(protein)
}
