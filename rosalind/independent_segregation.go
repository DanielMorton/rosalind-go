package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input from file
	file, err := os.Open("rosalind_indc.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	nStr := strings.TrimSpace(scanner.Text())
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	results := calculateProbabilities(n)

	// Print results
	for _, v := range results {
		fmt.Printf("%.3f ", v)
	}
	fmt.Println()
}

func calculateProbabilities(n int) []float64 {
	results := make([]float64, 2*n)
	logFact := make([]float64, 2*n+1)

	// Precompute log factorials
	logFact[0] = 0
	for i := 1; i <= 2*n; i++ {
		logFact[i] = logFact[i-1] + math.Log(float64(i))
	}

	log2 := math.Log(2)
	for k := 1; k <= 2*n; k++ {
		sum := 0.0
		for i := k; i <= 2*n; i++ {
			// Log of binomial coefficient
			logCoeff := logFact[2*n] - logFact[i] - logFact[2*n-i]
			// Log of probability
			logProb := logCoeff - float64(2*n)*log2
			sum += math.Exp(logProb)
		}
		results[k-1] = math.Log10(sum)
	}

	return results
}
