package main

import (
	"fmt"
	"os"
	"strings"
)

// Read input from file and parse the hidden path and transition matrix
func readInput(filename string) (string, map[string]map[string]float64) {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	path := lines[0]
	transition := make(map[string]map[string]float64)

	// Parse transition matrix header
	states := strings.Fields(lines[2]) // Extract states from the header

	// Initialize transition map for each state
	for _, state := range states {
		transition[state] = make(map[string]float64)
	}

	// Parse transition probabilities
	for i := 3; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		parts := strings.Fields(lines[i])
		currentState := parts[0]
		for j, prob := range parts[1:] {
			transition[currentState][states[j]] = parseFloat(prob)
		}
	}

	return path, transition
}

// Convert string to float64
func parseFloat(value string) float64 {
	var f float64
	fmt.Sscanf(value, "%f", &f)
	return f
}

// Calculate the probability of the hidden path
func calculatePathProbability(path string, transition map[string]map[string]float64) float64 {
	probability := 1.0
	numStates := len(path)

	// Since the initial state probabilities are equal, we divide by the number of states
	probability *= 1.0 / float64(len(transition))

	// Now multiply by the transition probabilities for each step in the path
	for i := 0; i < numStates-1; i++ {
		currentState := string(path[i])
		nextState := string(path[i+1])
		probability *= transition[currentState][nextState]
	}

	return probability
}

func main() {
	filename := "rosalind_ba10a.txt"
	path, transition := readInput(filename)

	probability := calculatePathProbability(path, transition)
	fmt.Printf("%.11e\n", probability) // Print with higher precision
}
