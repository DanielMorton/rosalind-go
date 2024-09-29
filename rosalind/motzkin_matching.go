package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Modulo constant
const MOD = 1000000

// Memoization map
var memo map[string]int

// Function to check if two bases can pair
func canPair(a, b byte) bool {
	return (a == 'A' && b == 'U') || (a == 'U' && b == 'A') || (a == 'C' && b == 'G') || (a == 'G' && b == 'C')
}

// Function to calculate noncrossing matchings using dynamic programming (Motzkin numbers for RNA)
func motzkinMatchings(rna string) int {
	// If the string is empty, there's exactly 1 way (the empty matching)
	if len(rna) == 0 {
		return 1
	}

	// Check if the result for this substring has already been calculated
	if val, found := memo[rna]; found {
		return val
	}

	// Calculate the total number of matchings
	totalMatchings := motzkinMatchings(rna[1:]) // Case where first base is unmatched

	// Try pairing the first base with some valid base j
	n := len(rna)
	for j := 1; j < n; j++ {
		if canPair(rna[0], rna[j]) {
			// Recursively calculate the number of matchings inside and outside the pair (0, j)
			leftMatchings := motzkinMatchings(rna[1:j])
			rightMatchings := motzkinMatchings(rna[j+1:])
			// Multiply the matchings of the two halves and add to the total
			totalMatchings += (leftMatchings * rightMatchings) % MOD
			totalMatchings %= MOD
		}
	}

	// Store the result in the memoization map
	memo[rna] = totalMatchings
	return totalMatchings
}

func main() {
	// Read input from file
	data, err := ioutil.ReadFile("rosalind_motz.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the RNA sequence (ignore the label line starting with '>')
	lines := strings.Split(string(data), "\n")
	var rna string
	for _, line := range lines {
		if !strings.HasPrefix(line, ">") {
			rna += line
		}
	}

	// Initialize memoization map
	memo = make(map[string]int)

	// Calculate the total number of noncrossing matchings, modulo 1,000,000
	result := motzkinMatchings(rna)

	// Print the result
	fmt.Println(result)
}
