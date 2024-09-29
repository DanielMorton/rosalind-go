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
	file, err := os.Open("rosalind_foun.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstLine := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(firstLine[0])
	m, _ := strconv.Atoi(firstLine[1])

	scanner.Scan()
	secondLine := strings.Fields(scanner.Text())
	A := make([]int, len(secondLine))
	for i, v := range secondLine {
		A[i], _ = strconv.Atoi(v)
	}

	// Calculate probabilities
	result := calculateProbabilities(N, m, A)

	// Print the result
	for _, row := range result {
		for _, val := range row {
			fmt.Printf("%.12f ", val)
		}
		fmt.Println()
	}
}

func calculateProbabilities(N, m int, A []int) [][]float64 {
	result := make([][]float64, m)
	for i := range result {
		result[i] = make([]float64, len(A))
	}

	for j, a := range A {
		q := 1.0 - float64(a)/(2*float64(N))
		for i := 0; i < m; i++ {
			probability := math.Pow(q, float64(2*N))
			if probability > 0 {
				result[i][j] = math.Log10(probability)
			} else {
				result[i][j] = math.Inf(-1)
			}
			q = math.Pow(q, 2)
		}
	}

	return result
}
