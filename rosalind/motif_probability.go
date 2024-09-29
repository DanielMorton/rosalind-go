package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateProbability(n int, x float64, s string) float64 {
	// Calculate the probability of the string occurring once
	pString := 1.0
	for _, nucleotide := range s {
		switch nucleotide {
		case 'A', 'T':
			pString *= (1 - x) / 2
		case 'C', 'G':
			pString *= x / 2
		}
	}

	// Calculate the probability of the string not occurring in one trial
	pNotString := 1 - pString

	// Calculate the probability of the string not occurring in n trials
	pNotStringN := math.Pow(pNotString, float64(n))

	// Return the probability of the string occurring at least once
	return 1 - pNotStringN
}

func main() {
	// Read input from file
	file, err := os.Open("rosalind_rstr.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the first line
	scanner.Scan()
	firstLine := scanner.Text()
	parts := strings.Fields(firstLine)

	if len(parts) != 2 {
		fmt.Println("Invalid input format in the first line")
		return
	}

	n, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println("Error parsing N:", err)
		return
	}

	x, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		fmt.Println("Error parsing x:", err)
		return
	}

	// Read the second line
	scanner.Scan()
	s := strings.TrimSpace(scanner.Text())

	if s == "" {
		fmt.Println("Error: DNA string is empty")
		return
	}

	// Calculate the probability
	probability := calculateProbability(n, x, s)

	// Print the result rounded to 3 decimal places
	fmt.Printf("%.3f\n", probability)
}
