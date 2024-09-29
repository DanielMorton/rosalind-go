package main

import (
	"fmt"
	"os"
	"strings"
)

// Read input from file and parse the string, alphabet, hidden path, and emission matrix
func readInput(filename string) (string, []string, string, map[string]map[string]float64) {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	x := lines[0]                        // The emitted string
	alphabet := strings.Fields(lines[2]) // Alphabet (symbols)
	hiddenPath := lines[4]               // The hidden path π

	// Initialize emission map for each state
	emission := make(map[string]map[string]float64)

	// Parse emission probabilities
	for i := 7; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		parts := strings.Fields(lines[i])
		currentState := parts[0]

		// Ensure emission[currentState] is initialized before assignment
		if _, ok := emission[currentState]; !ok {
			emission[currentState] = make(map[string]float64)
		}

		for j, prob := range parts[1:] {
			emission[currentState][alphabet[j]] = parseFloat(prob)
		}
	}

	return x, alphabet, hiddenPath, emission
}

// Convert string to float64
func parseFloat(value string) float64 {
	var f float64
	fmt.Sscanf(value, "%f", &f)
	return f
}

// Calculate the conditional probability Pr(x|π)
func calculateConditionalProbability(x string, pi string, emission map[string]map[string]float64) float64 {
	probability := 1.0

	// Iterate over each character in x and the corresponding state in π
	for i := 0; i < len(x); i++ {
		symbol := string(x[i])        // Current symbol from x
		currentState := string(pi[i]) // Corresponding state from π
		probability *= emission[currentState][symbol]
	}

	return probability
}

// Main function
func main() {
	filename := "rosalind_ba10b.txt"
	x, _, pi, emission := readInput(filename)

	probability := calculateConditionalProbability(x, pi, emission)
	fmt.Printf("%.11e\n", probability)
}
