package main

import (
	"bufio"
	"fmt"
	"log"
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
	file, err := os.Open("rosalind_ba10e.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read threshold
	scanner.Scan()
	threshold, _ := strconv.ParseFloat(scanner.Text(), 64)

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
	hmm := buildProfileHMM(alignment, alphabet, threshold)

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

	// Print separator
	fmt.Println("--------")

	// Print emission probabilities
	fmt.Printf("\t%s\n", strings.Join(alphabet, "\t"))
	for _, state := range hmm.states {
		fmt.Printf("%s", state)
		for _, symbol := range alphabet {
			fmt.Printf("\t%.3f", hmm.emissionProb[state][symbol])
		}
		fmt.Println()
	}
}

func buildProfileHMM(alignment []string, alphabet []string, threshold float64) HMM {
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

	// Initialize transition and emission probabilities
	transitionProb := make(map[string]map[string]float64)
	emissionProb := make(map[string]map[string]float64)
	for _, state := range states {
		transitionProb[state] = make(map[string]float64)
		emissionProb[state] = make(map[string]float64)
		for _, toState := range states {
			transitionProb[state][toState] = 0
		}
		for _, symbol := range alphabet {
			emissionProb[state][symbol] = 0
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
					transitionProb[prevState][currState]++
					prevState = currState
				} else {
					currState := fmt.Sprintf("M%d", matchIndex)
					transitionProb[prevState][currState]++
					emissionProb[currState][string(seq[j])]++
					prevState = currState
				}
			} else if seq[j] != '-' {
				currState := fmt.Sprintf("I%d", matchIndex)
				transitionProb[prevState][currState]++
				emissionProb[currState][string(seq[j])]++
				prevState = currState
			}
		}
		transitionProb[prevState]["E"]++
	}

	// Normalize probabilities
	for _, state := range states {
		totalTrans := 0.0
		for _, count := range transitionProb[state] {
			totalTrans += count
		}
		if totalTrans > 0 {
			for toState := range transitionProb[state] {
				transitionProb[state][toState] /= totalTrans
			}
		}

		totalEmit := 0.0
		for _, count := range emissionProb[state] {
			totalEmit += count
		}
		if totalEmit > 0 {
			for symbol := range emissionProb[state] {
				emissionProb[state][symbol] /= totalEmit
			}
		}
	}

	return HMM{states, transitionProb, emissionProb}
}
