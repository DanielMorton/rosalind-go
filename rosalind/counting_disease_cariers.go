package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input from file
	content, err := os.ReadFile("rosalind_afrq.txt")
	if err != nil {
		panic(err)
	}

	// Split input into array of strings
	inputStr := strings.Fields(string(content))

	// Initialize a slice to store the results
	results := make([]float64, len(inputStr))

	// Compute the probability B[k] for each A[k]
	for i, val := range inputStr {
		a, err := strconv.ParseFloat(val, 64)
		if err != nil {
			panic(err)
		}
		q := math.Sqrt(a)
		results[i] = 1 - math.Pow(1-q, 2)
	}

	// Print the results
	for _, res := range results {
		fmt.Printf("%.3f ", res)
	}
}
