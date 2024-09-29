package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Function to calculate the probability of matching the string s given GC-content p
func matchProbability(s string, gcContent float64) float64 {
	prob := 1.0
	for _, char := range s {
		switch char {
		case 'G', 'C':
			prob *= gcContent / 2
		case 'A', 'T':
			prob *= (1 - gcContent) / 2
		}
	}
	return prob
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_eval.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the length of the DNA string n
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// Read the DNA string s
	scanner.Scan()
	s := scanner.Text()

	// Read the array of GC-content probabilities
	scanner.Scan()
	gcContentStrings := strings.Fields(scanner.Text())
	var gcContents []float64
	for _, gc := range gcContentStrings {
		p, _ := strconv.ParseFloat(gc, 64)
		gcContents = append(gcContents, p)
	}

	// Calculate the expected number of occurrences for each GC-content
	sLen := len(s)
	results := make([]float64, len(gcContents))

	for i, gcContent := range gcContents {
		prob := matchProbability(s, gcContent)
		expectedOccurrences := prob * float64(n-sLen+1)
		results[i] = expectedOccurrences
	}

	// Print the results
	for _, res := range results {
		fmt.Printf("%.3f ", math.Round(res*1000)/1000)
	}
	fmt.Println()
}
