package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input from the file "rosalind_ebin.txt"
	content, err := os.ReadFile("rosalind_ebin.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	n, _ := strconv.Atoi(lines[0])            // Read the integer n
	probabilities := strings.Fields(lines[1]) // Read the probabilities

	// Prepare an array to hold expected values
	expectedValues := make([]float64, len(probabilities))

	// Calculate expected values
	for i, prob := range probabilities {
		p, _ := strconv.ParseFloat(prob, 64) // Convert to float64
		expectedValues[i] = float64(n) * p   // E(Bin(n, p)) = n * p
	}

	// Print the results with one decimal place
	for _, value := range expectedValues {
		fmt.Printf("%f ", value)
	}
	fmt.Println() // New line at the end
}
