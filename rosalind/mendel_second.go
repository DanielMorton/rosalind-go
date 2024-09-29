package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

// Function to calculate binomial probability iteratively
func binomialProbability(n, k int, p float64) float64 {
	prob := 0.0

	for i := k; i <= n; i++ {
		coeff := 1.0
		for j := 0; j < i; j++ {
			coeff *= float64(n-j) / float64(j+1)
		}
		prob += coeff * math.Pow(p, float64(i)) * math.Pow(1-p, float64(n-i))
	}

	return prob
}

// Function to calculate the probability of at least N Aa Bb organisms
func calculateProbability(k, N int) float64 {
	totalOffspring := int(math.Pow(2, float64(k))) // 2^k offspring in the k-th generation
	pAaBb := 0.25                                  // Probability of being Aa Bb

	// Calculate the probability of at least N Aa Bb organisms
	prob := binomialProbability(totalOffspring, N, pAaBb)

	return prob
}

func main() {
	// Read input from file
	data, err := ioutil.ReadFile("rosalind_lia.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the input values
	input := strings.Fields(string(data))
	k, err1 := strconv.Atoi(input[0])
	N, err2 := strconv.Atoi(input[1])
	if err1 != nil || err2 != nil {
		log.Fatal("Error parsing input")
	}

	// Calculate the probability
	result := calculateProbability(k, N)

	// Print the result rounded to 3 decimal places
	fmt.Printf("%.3f\n", result)
}
