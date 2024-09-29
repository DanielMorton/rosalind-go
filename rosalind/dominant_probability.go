package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to calculate the probability of producing a dominant phenotype
func dominantPhenotypeProbability(k, m, n int) float64 {
	// Total number of organisms
	total := k + m + n

	// Total number of possible pairs
	totalPairs := float64(total * (total - 1) / 2)

	// Recessive-recessive pairs (n, n): 100% recessive offspring
	recessiveRecessive := float64(n * (n - 1) / 2)

	// Heterozygous-recessive pairs (m, n): 50% chance of recessive offspring
	heteroRecessive := float64(m * n)

	// Heterozygous-heterozygous pairs (m, m): 25% chance of recessive offspring
	heteroHetero := float64(m * (m - 1) / 2)

	// Total probability of recessive offspring
	recessiveProb := (recessiveRecessive + 0.5*heteroRecessive + 0.25*heteroHetero) / totalPairs

	// Probability of producing dominant phenotype
	dominantProb := 1 - recessiveProb

	return dominantProb
}

func main() {
	// Open the file "rosalind_iprb.txt"
	file, err := os.Open("rosalind_iprb.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file and extract k, m, n
	scanner := bufio.NewScanner(file)
	scanner.Scan() // Read the first line

	// Split the line into separate components
	line := scanner.Text()
	values := strings.Split(line, " ")

	// Convert strings to integers
	k, _ := strconv.Atoi(values[0])
	m, _ := strconv.Atoi(values[1])
	n, _ := strconv.Atoi(values[2])

	// Calculate the probability of producing dominant offspring
	result := dominantPhenotypeProbability(k, m, n)

	// Print the result
	fmt.Printf("Probability of dominant phenotype: %.5f\n", result)
}
