package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input from file
	file, err := os.Open("rosalind_ba10c.txt")
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

	// Viterbi algorithm
	T := len(x)
	viterbi := make([]map[string]float64, T)
	backpointer := make([]map[string]string, T)

	// Initialization
	viterbi[0] = make(map[string]float64)
	backpointer[0] = make(map[string]string)
	for _, state := range states {
		viterbi[0][state] = math.Log(1.0/float64(len(states))) + math.Log(emissionMatrix[state][string(x[0])])
		backpointer[0][state] = ""
	}

	// Recursion
	for t := 1; t < T; t++ {
		viterbi[t] = make(map[string]float64)
		backpointer[t] = make(map[string]string)
		for _, state := range states {
			maxProb := math.Inf(-1)
			maxState := ""
			for _, prevState := range states {
				prob := viterbi[t-1][prevState] + math.Log(transitionMatrix[prevState][state]) + math.Log(emissionMatrix[state][string(x[t])])
				if prob > maxProb {
					maxProb = prob
					maxState = prevState
				}
			}
			viterbi[t][state] = maxProb
			backpointer[t][state] = maxState
		}
	}

	// Termination
	maxProb := math.Inf(-1)
	maxState := ""
	for _, state := range states {
		if viterbi[T-1][state] > maxProb {
			maxProb = viterbi[T-1][state]
			maxState = state
		}
	}

	// Backtracking
	path := make([]string, T)
	path[T-1] = maxState
	for t := T - 2; t >= 0; t-- {
		path[t] = backpointer[t+1][path[t+1]]
	}

	// Output result
	fmt.Println(strings.Join(path, ""))
}
