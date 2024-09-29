package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Function to read input data from the file
func readInput(filename string) (string, []string, []string, [][]float64, [][]float64) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the string x
	scanner.Scan()
	x := scanner.Text()

	// Read the alphabet Σ
	scanner.Scan()
	alphabet := strings.Fields(scanner.Text())

	// Read the states
	scanner.Scan()
	states := strings.Fields(scanner.Text())

	// Skip the '--------'
	scanner.Scan()

	// Read the transition matrix
	transition := make([][]float64, len(states))
	for i := range states {
		scanner.Scan()
		row := strings.Fields(scanner.Text())[1:] // Skip the state label
		transition[i] = make([]float64, len(states))
		for j := range row {
			transition[i][j], _ = strconv.ParseFloat(row[j], 64)
		}
	}

	// Skip the '--------'
	scanner.Scan()

	// Read the emission matrix
	emission := make([][]float64, len(states))
	for i := range states {
		scanner.Scan()
		row := strings.Fields(scanner.Text())[1:] // Skip the state label
		emission[i] = make([]float64, len(alphabet))
		fmt.Println(row)
		for j := range row {
			emission[i][j], _ = strconv.ParseFloat(row[j], 64)
		}
	}

	return x, alphabet, states, transition, emission
}

// Helper function to find index in a slice
func indexOf(s string, arr []string) int {
	for i, val := range arr {
		if val == s {
			return i
		}
	}
	return -1
}

// Viterbi algorithm implementation
func viterbi(x string, alphabet []string, states []string, transition [][]float64, emission [][]float64) string {
	n := len(x)
	k := len(states)

	// Initialize dynamic programming table and path table
	dp := make([][]float64, k)
	path := make([][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]float64, n)
		path[i] = make([]int, n)
	}

	// Initialize base cases
	for i := 0; i < k; i++ {
		emissionIndex := indexOf(string(x[0]), alphabet)
		dp[i][0] = emission[i][emissionIndex]
		path[i][0] = 0
	}

	// Fill the DP table
	for t := 1; t < n; t++ {
		emissionIndex := indexOf(string(x[t]), alphabet)
		for i := 0; i < k; i++ {
			maxProb := -1.0
			maxState := -1
			for j := 0; j < k; j++ {
				prob := dp[j][t-1] * transition[j][i] * emission[i][emissionIndex]
				if prob > maxProb {
					maxProb = prob
					maxState = j
				}
			}
			dp[i][t] = maxProb
			path[i][t] = maxState
		}
	}

	// Find the ending state with the highest probability
	maxFinalProb := -1.0
	lastState := -1
	for i := 0; i < k; i++ {
		if dp[i][n-1] > maxFinalProb {
			maxFinalProb = dp[i][n-1]
			lastState = i
		}
	}

	// Backtrack to find the best path
	bestPath := make([]int, n)
	bestPath[n-1] = lastState
	for t := n - 2; t >= 0; t-- {
		bestPath[t] = path[bestPath[t+1]][t+1]
	}

	// Convert the state indices to state labels
	result := make([]string, n)
	for t := 0; t < n; t++ {
		result[t] = states[bestPath[t]]
	}

	return strings.Join(result, "")
}

func main() {
	// Read the input
	x, alphabet, states, transition, emission := readInput("rosalind_ba10c.txt")

	fmt.Println(transition)
	fmt.Println(emission)
	// Run the Viterbi algorithm
	result := viterbi(x, alphabet, states, transition, emission)

	// Print the result
	fmt.Println(result)
}
