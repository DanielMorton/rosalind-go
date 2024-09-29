package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var masses = map[byte]float64{
	'A': 71.03711,
	'C': 103.00919,
	'D': 115.02694,
	'E': 129.04259,
	'F': 147.06841,
	'G': 57.02146,
	'H': 137.05891,
	'I': 113.08406,
	'K': 128.09496,
	'L': 113.08406,
	'M': 131.04049,
	'N': 114.04293,
	'P': 97.05276,
	'Q': 128.05858,
	'R': 156.10111,
	'S': 87.03203,
	'T': 101.04768,
	'V': 99.06841,
	'W': 186.07931,
	'Y': 163.06333,
}

func calculateSpectrum(s string) []float64 {
	n := len(s)
	spectrum := make([]float64, 2*n)

	for i := 1; i <= n; i++ {
		prefix := s[:i]
		suffix := s[n-i:]

		prefixMass := 0.0
		suffixMass := 0.0

		for j := 0; j < i; j++ {
			prefixMass += masses[prefix[j]]
			suffixMass += masses[suffix[j]]
		}

		spectrum[i-1] = prefixMass
		spectrum[n+i-1] = suffixMass
	}

	return spectrum
}

func multisetDifference(a, b []float64) int {
	count := 0
	tolerance := 0.001

	for _, x := range a {
		for j, y := range b {
			if math.Abs(x-y) < tolerance {
				count++
				b[j] = math.Inf(1) // Mark as used
				break
			}
		}
	}

	return count
}

func main() {
	file, err := os.Open("rosalind_prsm.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	proteins := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		proteins[i] = scanner.Text()
	}

	var R []float64
	for scanner.Scan() {
		mass, _ := strconv.ParseFloat(scanner.Text(), 64)
		R = append(R, mass)
	}

	maxMultiplicity := 0
	var bestProtein string

	for _, protein := range proteins {
		spectrum := calculateSpectrum(protein)
		multiplicity := multisetDifference(R, spectrum)

		if multiplicity > maxMultiplicity {
			maxMultiplicity = multiplicity
			bestProtein = protein
		}
	}

	fmt.Println(maxMultiplicity)
	fmt.Println(bestProtein)
}
