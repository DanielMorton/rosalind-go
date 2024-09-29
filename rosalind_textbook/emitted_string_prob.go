package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input from file
	file, err := os.Open("rosalind_ba10d.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read emitted string
	scanner.Scan()
	x := scanner.Text()

	// Skip the separator line
	scanner.Scan()

	// Read alphabet
	scanner.Scan()
	alphabet := strings.Fields(scanner.Text())

	// Skip the separator line
	scanner.Scan()

	// Read states
	scanner.Scan()
	states := strings.Fields(scanner.Text())

	// Skip the separator lines
	scanner.Scan()
	scanner.Scan()

	// Initialize transition matrix
	transitionMatrix := make(map[string]map[string]float64)
	for _, state := range states {
		transitionMatrix[state] = make(map[string]float64)
	}

	// Read transition matrix
	for _, fromState := range states {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		for i, prob := range line[1:] {
			p, _ := strconv.ParseFloat(prob, 64)
			transitionMatrix[fromState][states[i]] = p
		}
	}

	// Skip the separator line
	scanner.Scan()

	// Initialize emission matrix
	emissionMatrix := make(map[string]map[string]float64)
	for _, state := range states {
		emissionMatrix[state] = make(map[string]float64)
	}

	// Read emission matrix
	scanner.Scan() // Skip the header line
	for _, state := range states {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		for i, prob := range line[1:] {
			p, _ := strconv.ParseFloat(prob, 64)
			emissionMatrix[state][alphabet[i]] = p
		}
	}

	// Forward algorithm
	T := len(x)
	forward := make([]map[string]float64, T)

	// Initialization
	forward[0] = make(map[string]float64)
	for _, state := range states {
		forward[0][state] = 1.0 / float64(len(states)) * emissionMatrix[state][string(x[0])]
	}

	// Recursion
	for t := 1; t < T; t++ {
		forward[t] = make(map[string]float64)
		for _, state := range states {
			sum := 0.0
			for _, prevState := range states {
				sum += forward[t-1][prevState] * transitionMatrix[prevState][state]
			}
			forward[t][state] = sum * emissionMatrix[state][string(x[t])]
		}
	}

	// Termination
	probability := 0.0
	for _, state := range states {
		probability += forward[T-1][state]
	}

	// Output result
	fmt.Printf("%.14e\n", probability)
}
