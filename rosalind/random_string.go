package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// calculateLogProbability calculates the log probability of matching the sequence given the GC-content
func calculateLogProbability(sequence string, gcContent float64) float64 {
	// Calculate the probabilities of 'G'/'C' and 'A'/'T'
	pGC := gcContent / 2
	pAT := (1 - gcContent) / 2

	// Calculate the log probability of matching the sequence
	var logProb float64 = 0.0
	for _, char := range sequence {
		switch char {
		case 'A', 'T':
			logProb += math.Log10(pAT)
		case 'C', 'G':
			logProb += math.Log10(pGC)
		default:
			return math.Inf(-1) // Return negative infinity for invalid characters
		}
	}

	return logProb
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_prob.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the DNA string
	scanner.Scan()
	sequence := scanner.Text()

	// Read the GC-content values
	scanner.Scan()
	gcContentStrings := strings.Fields(scanner.Text())

	// Convert GC-content values to floats
	var gcContents []float64
	for _, str := range gcContentStrings {
		gcContent, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("Error parsing GC-content value:", err)
			return
		}
		gcContents = append(gcContents, gcContent)
	}

	// Calculate the logarithm of the probability for each GC-content value
	var results []float64
	for _, gcContent := range gcContents {
		logProbability := calculateLogProbability(sequence, gcContent)
		results = append(results, logProbability)
	}

	// Print the results
	for _, result := range results {
		fmt.Printf("%.3f ", result)
	}
	fmt.Println()
}
