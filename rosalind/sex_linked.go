package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input from the file "rosalind_sexl.txt"
	content, err := os.ReadFile("rosalind_sexl.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	A := strings.Fields(lines[0]) // Read the proportions of males with recessive genes
	n := len(A)
	B := make([]float64, n)

	// Calculate the probabilities for females being carriers
	for k := 0; k < n; k++ {
		proportion, _ := strconv.ParseFloat(A[k], 64)
		B[k] = 2 * proportion * (1 - proportion) // Calculate the carrier probability
	}

	// Print the results with appropriate formatting
	for _, value := range B {
		fmt.Printf("%.10f ", value) // Print with 10 decimal places
	}
	fmt.Println() // New line at the end
}
