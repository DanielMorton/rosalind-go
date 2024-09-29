package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to compute the expected number of dominant phenotype offspring
func expectedDominantOffspring(couples [6]int) float64 {
	// Probabilities of dominant phenotype for each genotype pairing
	probs := [6]float64{
		1.0,  // AA-AA
		1.0,  // AA-Aa
		1.0,  // AA-aa
		0.75, // Aa-Aa
		0.50, // Aa-aa
		0.0,  // aa-aa
	}

	expected := 0.0
	for i, numCouples := range couples {
		expected += 2 * probs[i] * float64(numCouples)
	}

	return expected
}

func main() {
	// Open the file "rosalind_iprb.txt"
	file, err := os.Open("rosalind_iev.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	// Split the line into integers
	parts := strings.Fields(line)
	var couples [6]int
	for i, part := range parts {
		couples[i], err = strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error parsing integer:", err)
			return
		}
	}

	// Calculate the expected number of dominant phenotype offspring
	expected := expectedDominantOffspring(couples)

	// Print the result
	fmt.Printf("%.5f\n", expected)
}
