package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var aminoAcidMasses = []int{57, 71, 87, 97, 99, 101, 103, 113, 114, 115, 128, 129, 131, 137, 147, 156, 163, 186}

type Peptide []int

func expand(peptides []Peptide) []Peptide {
	var expanded []Peptide
	for _, peptide := range peptides {
		for _, mass := range aminoAcidMasses {
			newPeptide := append(Peptide(nil), peptide...)
			newPeptide = append(newPeptide, mass)
			expanded = append(expanded, newPeptide)
		}
	}
	return expanded
}

func mass(peptide Peptide) int {
	sum := 0
	for _, mass := range peptide {
		sum += mass
	}
	return sum
}

func cyclospectrum(peptide Peptide) []int {
	n := len(peptide)
	extendedPeptide := append(peptide, peptide...)
	spectrum := []int{0, mass(peptide)} // Include 0 and total mass

	for length := 1; length < n; length++ {
		for i := 0; i < n; i++ {
			subPeptide := extendedPeptide[i : i+length]
			spectrum = append(spectrum, mass(subPeptide))
		}
	}

	sort.Ints(spectrum)
	return spectrum
}

func isConsistent(peptide Peptide, spectrum []int) bool {
	peptideSpectrum := linearSpectrum(peptide)
	i, j := 0, 0
	for i < len(peptideSpectrum) && j < len(spectrum) {
		if peptideSpectrum[i] < spectrum[j] {
			return false
		} else if peptideSpectrum[i] > spectrum[j] {
			j++
		} else {
			i++
			j++
		}
	}
	return true
}

func linearSpectrum(peptide Peptide) []int {
	spectrum := []int{0}
	for i := 0; i < len(peptide); i++ {
		for j := i + 1; j <= len(peptide); j++ {
			subPeptide := peptide[i:j]
			spectrum = append(spectrum, mass(subPeptide))
		}
	}
	sort.Ints(spectrum)
	return spectrum
}

func cyclopeptideSequencing(spectrum []int) []Peptide {
	var results []Peptide
	peptides := []Peptide{{}} // Start with an empty peptide
	parentMass := spectrum[len(spectrum)-1]

	for len(peptides) > 0 {
		peptides = expand(peptides)
		var newPeptides []Peptide
		for _, peptide := range peptides {
			if mass(peptide) == parentMass {
				if equal(cyclospectrum(peptide), spectrum) {
					results = append(results, peptide)
				}
			} else if mass(peptide) < parentMass && isConsistent(peptide, spectrum) {
				newPeptides = append(newPeptides, peptide)
			}
		}
		peptides = newPeptides
	}

	return results
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("rosalind_ba4e.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	spectrumStr := strings.Fields(scanner.Text())

	spectrum := make([]int, len(spectrumStr))
	for i, s := range spectrumStr {
		spectrum[i], _ = strconv.Atoi(s)
	}

	results := cyclopeptideSequencing(spectrum)

	for i, peptide := range results {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(peptide)), "-"), "[]"))
	}
	fmt.Println()
}
