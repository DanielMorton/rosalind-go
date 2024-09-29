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
	// Read input from file
	file, err := os.Open("rosalind_ba10g.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read text
	scanner.Scan()
	text := scanner.Text()

	// Skip separator line
	scanner.Scan()

	// Read threshold and pseudocount
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	threshold, _ := strconv.ParseFloat(params[0], 64)
	pseudocount, _ := strconv.ParseFloat(params[1], 64)

	// Skip separator line
	scanner.Scan()

	// Read alphabet
	scanner.Scan()
	alphabet := strings.Fields(scanner.Text())

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

	// Build profile HMM
	hmm := buildProfileHMM(alignment, alphabet, threshold, pseudocount)

	// Print transition probabilities
	fmt.Print("\t")
	fmt.Println(strings.Join(hmm.states, "\t"))
	for _, fromState := range hmm.states {
		fmt.Printf("%s", fromState)
		for _, toState := range hmm.states {
			fmt.Printf("\t%.3f", hmm.transitionProb[fromState][toState])
		}
		fmt.Println()
	}

	// Print input data for verification
	fmt.Println("Input Text:", text)
	fmt.Println("Threshold:", threshold)
	fmt.Println("Pseudocount:", pseudocount)
	fmt.Println("Alphabet:", alphabet)
	fmt.Println("Alignment:", alignment)

	// Print HMM information for debugging
	fmt.Println("States:", hmm.states)
	fmt.Println("Transition Probabilities:")
	for fromState, toStates := range hmm.transitionProb {
		fmt.Printf("%s: %v\n", fromState, toStates)
	}
	fmt.Println("Emission Probabilities:")
	for state, emissions := range hmm.emissionProb {
		fmt.Printf("%s: %v\n", state, emissions)
	}

	// Find optimal hidden path
	optimalPath, viterbiMatrix, backpointer, err := viterbiAlgorithm(text, hmm)
	if err != nil {
		log.Fatalf("Error in Viterbi algorithm: %v", err)
	}

	// Output result
	fmt.Println("Optimal Path:", strings.Join(optimalPath, " "))

	// Print Viterbi matrix and backpointer for debugging (only first and last few steps)
	fmt.Println("Viterbi Matrix and Backpointer (first 3 and last 3 steps):")
	for t := 0; t < 3; t++ {
		fmt.Printf("t=%d:\n", t)
		for state, prob := range viterbiMatrix[t] {
			if prob > math.Inf(-1) {
				fmt.Printf("  %s: %.4f (from: %s)\n", state, prob, backpointer[t][state])
			}
		}
	}
	fmt.Println("...")
	for t := len(viterbiMatrix) - 3; t < len(viterbiMatrix); t++ {
		fmt.Printf("t=%d:\n", t)
		for state, prob := range viterbiMatrix[t] {
			if prob > math.Inf(-1) {
				fmt.Printf("  %s: %.4f (from: %s)\n", state, prob, backpointer[t][state])
			}
		}
	}
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

	// Define possible transitions
	possibleTransitions := map[string][]string{
		"S":  {"I0", "M1", "D1"},
		"I0": {"I0", "M1", "D1"},
	}
	for i := 1; i <= numMatches; i++ {
		Mi := fmt.Sprintf("M%d", i)
		Di := fmt.Sprintf("D%d", i)
		Ii := fmt.Sprintf("I%d", i)
		if i < numMatches {
			possibleTransitions[Mi] = []string{fmt.Sprintf("M%d", i+1), fmt.Sprintf("D%d", i+1), Ii}
			possibleTransitions[Di] = []string{fmt.Sprintf("M%d", i+1), fmt.Sprintf("D%d", i+1), Ii}
			possibleTransitions[Ii] = []string{fmt.Sprintf("M%d", i+1), fmt.Sprintf("D%d", i+1), Ii}
		} else {
			possibleTransitions[Mi] = []string{"E", Ii}
			possibleTransitions[Di] = []string{"E", Ii}
			possibleTransitions[Ii] = []string{"E", Ii}
		}
	}
	fmt.Println(possibleTransitions)

	// Calculate initial probabilities, add pseudocounts, and renormalize
	transitionProb := make(map[string]map[string]float64)
	emissionProb := make(map[string]map[string]float64)
	for _, state := range states {
		transitionProb[state] = make(map[string]float64)
		emissionProb[state] = make(map[string]float64)

		// Initial transition probabilities
		totalTrans := 0.0
		for _, count := range transitionCount[state] {
			totalTrans += count
		}
		for toState, count := range transitionCount[state] {
			if totalTrans > 0 {
				transitionProb[state][toState] = count / totalTrans
			}
		}

		// Add pseudocounts and renormalize transitions
		totalTransWithPseudo := 0.0
		for _, toState := range possibleTransitions[state] {
			transitionProb[state][toState] += pseudocount
			totalTransWithPseudo += transitionProb[state][toState]
		}
		for toState := range transitionProb[state] {
			if totalTransWithPseudo > 0 {
				transitionProb[state][toState] /= totalTransWithPseudo
			}
		}

		// Initial emission probabilities (only for non-deletion states)
		if !strings.HasPrefix(state, "D") && state != "S" && state != "E" {
			totalEmit := 0.0
			for _, count := range emissionCount[state] {
				totalEmit += count
			}
			for symbol, count := range emissionCount[state] {
				if totalEmit > 0 {
					emissionProb[state][symbol] = count / totalEmit
				}
			}

			// Add pseudocounts and renormalize emissions
			totalEmitWithPseudo := 0.0
			for _, symbol := range alphabet {
				emissionProb[state][symbol] += pseudocount
				totalEmitWithPseudo += emissionProb[state][symbol]
			}
			for symbol := range emissionProb[state] {
				if totalEmitWithPseudo > 0 {
					emissionProb[state][symbol] /= totalEmitWithPseudo
				}
			}
		}
	}

	return HMM{states, transitionProb, emissionProb}
}

func viterbiAlgorithm(text string, hmm HMM) ([]string, []map[string]float64, []map[string]string, error) {
	T := len(text)
	viterbi := make([]map[string]float64, T+1)
	backpointer := make([]map[string]string, T+1)

	// Initialization
	viterbi[0] = make(map[string]float64)
	backpointer[0] = make(map[string]string)
	viterbi[0]["S"] = 0.0 // log(1) = 0
	for _, state := range hmm.states {
		if state != "S" {
			viterbi[0][state] = math.Inf(-1)
			backpointer[0][state] = ""
		}
	}

	// Print initial viterbi matrix
	fmt.Println("Initial Viterbi Matrix:")
	printViterbiMatrix(viterbi[0])

	// Handle the first position separately
	viterbi[1] = make(map[string]float64)
	backpointer[1] = make(map[string]string)
	for _, state := range hmm.states {
		if state == "S" || state == "E" {
			continue
		}
		if trans, ok := hmm.transitionProb["S"][state]; ok && trans > 0 {
			var emissionProb float64
			if strings.HasPrefix(state, "M") || strings.HasPrefix(state, "I") {
				emissionProb = hmm.emissionProb[state][string(text[0])]
			} else {
				emissionProb = 1.0
			}
			if emissionProb > 0 {
				viterbi[1][state] = math.Log(trans) + math.Log(emissionProb)
				backpointer[1][state] = "S"
			}
		}
	}

	fmt.Println("\nViterbi Matrix after position 1:")
	printViterbiMatrix(viterbi[1])

	// Recursion
	for t := 2; t <= T; t++ {
		fmt.Printf("\nProcessing position %d, symbol %s\n", t, string(text[t-1]))
		viterbi[t] = make(map[string]float64)
		backpointer[t] = make(map[string]string)
		for _, state := range hmm.states {
			if state == "S" || state == "E" {
				continue
			}
			maxProb := math.Inf(-1)
			maxState := ""
			fmt.Printf("  Considering state %s\n", state)
			for _, prevState := range hmm.states {
				if prevState == "S" || prevState == "E" {
					continue
				}
				if trans, ok := hmm.transitionProb[prevState][state]; ok && trans > 0 {
					var emissionProb float64
					if strings.HasPrefix(state, "M") || strings.HasPrefix(state, "I") {
						emissionProb = hmm.emissionProb[state][string(text[t-1])]
					} else {
						emissionProb = 1.0
					}
					if emissionProb > 0 {
						prob := viterbi[t-1][prevState] + math.Log(trans) + math.Log(emissionProb)
						fmt.Printf("    From %s: prob = %.4f (prev: %.4f, trans: %.4f, emit: %.4f)\n",
							prevState, prob, viterbi[t-1][prevState], math.Log(trans), math.Log(emissionProb))
						if prob > maxProb {
							maxProb = prob
							maxState = prevState
						}
					} else {
						fmt.Printf("    From %s: Skipped due to zero emission probability\n", prevState)
					}
				} else {
					fmt.Printf("    From %s: Skipped due to zero transition probability\n", prevState)
				}
			}
			if maxState != "" {
				viterbi[t][state] = maxProb
				backpointer[t][state] = maxState
				fmt.Printf("  Chose %s for %s with probability %.4f\n", maxState, state, maxProb)
			}
		}

		// Print viterbi matrix after processing this position
		fmt.Printf("\nViterbi Matrix after position %d:\n", t)
		printViterbiMatrix(viterbi[t])
	}

	// Termination
	maxProb := math.Inf(-1)
	maxState := ""
	for _, state := range hmm.states {
		if state != "S" && state != "E" {
			if trans, ok := hmm.transitionProb[state]["E"]; ok && trans > 0 {
				prob := viterbi[T][state] + math.Log(trans)
				if prob > maxProb {
					maxProb = prob
					maxState = state
				}
			}
		}
	}
	if maxState == "" {
		return nil, nil, nil, fmt.Errorf("no valid path to end state")
	}

	// Backtracking
	path := make([]string, T)
	path[T-1] = maxState
	for t := T - 2; t >= 0; t-- {
		path[t] = backpointer[t+2][path[t+1]]
		if path[t] == "" {
			return nil, nil, nil, fmt.Errorf("backtracking failed at position %d", t)
		}
	}

	return path, viterbi, backpointer, nil
}

func printViterbiMatrix(matrix map[string]float64) {
	for state, prob := range matrix {
		if prob > math.Inf(-1) {
			fmt.Printf("  %s: %.4f\n", state, prob)
		}
	}
}

func isValidTransition(from, to string) bool {
	fromType := from[0:1]
	toType := to[0:1]
	fromIndex, _ := strconv.Atoi(from[1:])
	toIndex, _ := strconv.Atoi(to[1:])

	if from == "S" {
		return toType == "M" || toType == "D" || toType == "I"
	}

	if fromType == "M" || fromType == "D" {
		return (toType == "M" && toIndex == fromIndex+1) ||
			(toType == "D" && toIndex == fromIndex+1) ||
			(toType == "I" && toIndex == fromIndex)
	}

	if fromType == "I" {
		return (toType == "M" && toIndex == fromIndex+1) ||
			(toType == "D" && toIndex == fromIndex+1) ||
			(toType == "I" && toIndex == fromIndex)
	}

	return false
}

func (hmm HMM) printProbabilities() {
	fmt.Println("Transition Probabilities:")
	for fromState, toStates := range hmm.transitionProb {
		fmt.Printf("%s:\n", fromState)
		for toState, prob := range toStates {
			fmt.Printf("  %s: %f\n", toState, prob)
		}
	}

	fmt.Println("\nEmission Probabilities:")
	for state, emissions := range hmm.emissionProb {
		fmt.Printf("%s:\n", state)
		for symbol, prob := range emissions {
			fmt.Printf("  %s: %f\n", symbol, prob)
		}
	}
}
