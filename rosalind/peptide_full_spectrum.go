package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Amino acid mass table (monoisotopic masses)
var aminoAcidMasses = map[string]float64{
	"G": 57.021464, "A": 71.037114, "S": 87.032028, "P": 97.052764,
	"V": 99.068414, "T": 101.04768, "C": 103.00919, "I": 113.08406,
	"L": 113.08406, "N": 114.04293, "D": 115.02694, "Q": 128.05858,
	"K": 128.09496, "E": 129.04259, "M": 131.04049, "H": 137.05891,
	"F": 147.06841, "R": 156.10111, "Y": 163.06333, "W": 186.07931,
}

// Tolerance for floating-point comparison
const tolerance = 0.01

// Helper function to compare two float values within a tolerance
func floatEquals(a, b float64) bool {
	return math.Abs(a-b) <= tolerance
}

// Function to parse the input file
func parseInput(filename string) (float64, []float64) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var parentMass float64
	var masses []float64

	if scanner.Scan() {
		parentMass, _ = strconv.ParseFloat(scanner.Text(), 64)
	}

	for scanner.Scan() {
		mass, _ := strconv.ParseFloat(scanner.Text(), 64)
		masses = append(masses, mass)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	return parentMass, masses
}

// Function to find the closest amino acid given a mass difference
func findAminoAcid(diff float64) (string, bool) {
	for aa, mass := range aminoAcidMasses {
		if floatEquals(diff, mass) {
			return aa, true
		}
	}
	return "", false
}

// Function to attempt mapping using combinations of amino acids
func findAminoAcidByCombination(diff float64) (string, bool) {
	// Try to find a combination of two amino acids that sum up to the difference
	for aa1, mass1 := range aminoAcidMasses {
		for aa2, mass2 := range aminoAcidMasses {
			if floatEquals(mass1+mass2, diff) {
				return aa1 + aa2, true
			}
		}
	}
	return "", false
}

// Main function to reconstruct the protein sequence
func reconstructProtein(parentMass float64, ionMasses []float64) string {
	sort.Float64s(ionMasses)
	n := len(ionMasses) / 2 // The length of the peptide is half the number of ions

	var proteinSeq []string

	for i := 0; i < n; i++ {
		// Calculate the difference between consecutive b-ions
		diff := ionMasses[i+1] - ionMasses[i]
		aa, found := findAminoAcid(diff)
		if found {
			proteinSeq = append(proteinSeq, aa)
		} else {
			// Try combinations of amino acids
			aaCombo, comboFound := findAminoAcidByCombination(diff)
			if comboFound {
				proteinSeq = append(proteinSeq, aaCombo)
			} else {
				// Log the error but continue processing
				fmt.Printf("Warning: could not find amino acid for mass difference: %f\n", diff)
			}
		}
	}

	return strings.Join(proteinSeq, "")
}

func main() {
	// Input file
	filename := "rosalind_full.txt"

	// Parse input
	parentMass, ionMasses := parseInput(filename)

	// Reconstruct the protein sequence
	proteinSeq := reconstructProtein(parentMass, ionMasses)

	// Output the result
	fmt.Println(proteinSeq)
}
