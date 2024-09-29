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
	file, err := os.Open("rosalind_ba10a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read hidden path
	scanner.Scan()
	path := scanner.Text()

	// Skip the separator line
	scanner.Scan()

	// Read states
	scanner.Scan()
	states := strings.Fields(scanner.Text())

	// Skip the separator line
	scanner.Scan()

	// Read transition matrix
	transitionMatrix := make(map[string]map[string]float64)
	for _, state := range states {
		transitionMatrix[state] = make(map[string]float64)
	}

	scanner.Scan() // Skip the header line
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		fromState := line[0]
		for i, prob := range line[1:] {
			p, _ := strconv.ParseFloat(prob, 64)
			transitionMatrix[fromState][states[i]] = p
		}
	}

	// Calculate probability
	initialProb := 1.0 / float64(len(states)) // Equal initial probabilities
	probability := initialProb

	for i := 1; i < len(path); i++ {
		fromState := string(path[i-1])
		toState := string(path[i])
		probability *= transitionMatrix[fromState][toState]
	}

	// Output result
	fmt.Printf("%.14e\n", probability)
}
