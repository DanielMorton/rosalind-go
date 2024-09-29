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

type HMM struct {
	states         []string
	transitionProb map[string]map[string]float64
	emissionProb   map[string]map[string]float64
}

func main() {
	file, err := os.Open("rosalind_ba10g.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read Text
	scanner.Scan()
	text := scanner.Text()
	fmt.Printf("Text: %s\n", text)

	// Skip separator line
	scanner.Scan()

	// Read threshold and pseudocount
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	threshold, _ := strconv.ParseFloat(params[0], 64)
	pseudocount, _ := strconv.ParseFloat(params[1], 64)
	fmt.Printf("Threshold: %f, Pseudocount: %f\n", threshold, pseudocount)

	// Skip separator line
	scanner.Scan()

	// Read alphabet
	scanner.Scan()
	alphabet := strings.Fields(scanner.Text())
	fmt.Printf("Alphabet: %v\n", alphabet)

	// Skip separator line
	scanner.Scan()

	// Read alignment
	var alignment []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		alignment = append(alignment, line)
	}
	fmt.Printf("Alignment:\n%s\n", strings.Join(alignment, "\n"))

	// Build profile HMM
	hmm := buildProfileHMM(alignment, alphabet, threshold, pseudocount)
	fmt.Printf("HMM States: %v\n", hmm.states)

	// Run Viterbi algorithm
	path := viterbi(text, hmm)

	// Print result
	fmt.Println("Viterbi Path:")
	fmt.Println(strings.Join(path, " "))
}

func buildProfileHMM(alignment []string, alphabet []string, threshold, pseudocount float64) HMM {
	numSequences := len(alignment)
	alignmentLength := len(alignment[0])

	// Determine match columns
	matchColumns := make([]bool, alignmentLength)
	for j := 0; j < alignmentLength; j++ {
		nonGapCount := 0
		for i := 0; i < numSequences; i++ {
			if alignment[i][j] != '-' {
				nonGapCount++
			}
		}
		matchColumns[j] = float64(nonGapCount)/float64(numSequences) > 1-threshold
	}

	// Build states
	var states []string
	states = append(states, "S", "I0")
	numMatches := 0
	for _, isMatch := range matchColumns {
		if isMatch {
			numMatches++
			states = append(states, fmt.Sprintf("M%d", numMatches), fmt.Sprintf("D%d", numMatches), fmt.Sprintf("I%d", numMatches))
		}
	}
	states = append(states, "E")

	// Initialize transition and emission counts
	transitionCount := make(map[string]map[string]float64)
	emissionCount := make(map[string]map[string]float64)
	for _, state := range states {
		transitionCount[state] = make(map[string]float64)
		emissionCount[state] = make(map[string]float64)
		for _, toState := range states {
			transitionCount[state][toState] = pseudocount
		}
		for _, symbol := range alphabet {
			emissionCount[state][symbol] = pseudocount
		}
	}

	// Count transitions and emissions
	for _, seq := range alignment {
		prevState := "S"
		matchIndex := 0
		for j := 0; j < alignmentLength; j++ {
			if matchColumns[j] {
				matchIndex++
				if seq[j] == '-' {
					currState := fmt.Sprintf("D%d", matchIndex)
					transitionCount[prevState][currState]++
					prevState = currState
				} else {
					currState := fmt.Sprintf("M%d", matchIndex)
					transitionCount[prevState][currState]++
					emissionCount[currState][string(seq[j])]++
					prevState = currState
				}
			} else if seq[j] != '-' {
				currState := fmt.Sprintf("I%d", matchIndex)
				transitionCount[prevState][currState]++
				emissionCount[currState][string(seq[j])]++
				prevState = currState
			}
		}
		transitionCount[prevState]["E"]++
	}

	// Calculate probabilities
	transitionProb := make(map[string]map[string]float64)
	emissionProb := make(map[string]map[string]float64)
	for _, state := range states {
		transitionProb[state] = make(map[string]float64)
		emissionProb[state] = make(map[string]float64)

		totalTrans := 0.0
		for _, count := range transitionCount[state] {
			totalTrans += count
		}
		for toState, count := range transitionCount[state] {
			transitionProb[state][toState] = count / totalTrans
		}

		totalEmit := 0.0
		for _, count := range emissionCount[state] {
			totalEmit += count
		}
		for symbol, count := range emissionCount[state] {
			emissionProb[state][symbol] = count / totalEmit
		}
	}

	return HMM{states, transitionProb, emissionProb}
}

func viterbi(text string, hmm HMM) []string {
	T := len(text)
	viterbi := make([]map[string]float64, T)
	backpointer := make([]map[string]string, T)

	// Initialization
	viterbi[0] = make(map[string]float64)
	backpointer[0] = make(map[string]string)
	for _, state := range hmm.states {
		viterbi[0][state] = math.Log(1.0/float64(len(hmm.states))) + math.Log(hmm.emissionProb[state][string(text[0])])
		backpointer[0][state] = ""
	}

	// Recursion
	for t := 1; t < T; t++ {
		viterbi[t] = make(map[string]float64)
		backpointer[t] = make(map[string]string)
		for _, state := range hmm.states {
			maxProb := math.Inf(-1)
			maxState := ""
			for _, prevState := range hmm.states {
				prob := viterbi[t-1][prevState] + math.Log(hmm.transitionProb[prevState][state]) + math.Log(hmm.emissionProb[state][string(text[t])])
				if prob >= maxProb {
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
	for _, state := range hmm.states {
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
	return path
}
