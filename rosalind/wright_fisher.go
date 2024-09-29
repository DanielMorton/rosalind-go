package main

import (
	"fmt"
	"math/rand"
	"os"
)

// Function to calculate the probability of having at least k recessive alleles
func calculateProbability(N, m, g, k int) float64 {
	successes := 0
	trials := 100000 // Number of trials for Monte Carlo simulation

	for i := 0; i < trials; i++ {
		// Start with m dominant alleles
		currentDominant := m
		for gen := 0; gen < g; gen++ {
			// Calculate the probability of choosing a dominant allele
			p := float64(currentDominant) / float64(2*N)
			// Calculate the number of dominant alleles in the next generation
			currentDominant = randBinomial(2*N, p)
		}

		// Calculate the number of recessive alleles
		currentRecessive := 2*N - currentDominant
		if currentRecessive >= k {
			successes++
		}
	}

	// Calculate the probability
	return float64(successes) / float64(trials)
}

// Helper function to simulate a binomial random variable
func randBinomial(n int, p float64) int {
	successes := 0
	for i := 0; i < n; i++ {
		if rand.Float64() < p {
			successes++
		}
	}
	return successes
}

func main() {
	// Read input from file "rosalind_wfmd.txt"
	content, err := os.ReadFile("rosalind_wfmd.txt")
	if err != nil {
		panic(err)
	}

	var N, m, g, k int
	fmt.Sscanf(string(content), "%d %d %d %d", &N, &m, &g, &k)

	// Calculate the probability
	probability := calculateProbability(N, m, g, k)

	// Print the result with three decimal places
	fmt.Printf("%.3f\n", probability)
}
