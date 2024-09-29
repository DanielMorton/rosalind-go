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
	file, err := os.Open("rosalind_ba10f.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read threshold and pseudocount
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	threshold, _ := strconv.ParseFloat(params[0], 64)
	pseudocount, _ := strconv.ParseFloat(params[1], 64)
	fmt.Println(pseudocount)

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

/*func buildProfileHMM(alignment []string, alphabet []string, threshold, pseudocount float64) HMM {
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
}*/

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
