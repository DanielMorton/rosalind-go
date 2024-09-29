package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	gapPenalty      = -1.0
	mismatchPenalty = -1.0
)

// Read and parse FASTA file
func readFasta(filename string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var sequences []string
	var currentSeq strings.Builder
	inSequence := false

	for _, line := range strings.Split(string(file), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ">") {
			if inSequence {
				sequences = append(sequences, currentSeq.String())
				currentSeq.Reset()
			}
			inSequence = true
		} else if inSequence {
			currentSeq.WriteString(line)
		}
	}

	if inSequence {
		sequences = append(sequences, currentSeq.String())
	}

	return sequences, nil
}

// Compute the alignment score for all possible alignments of multiple sequences
func alignMultipleSequences(sequences []string) (float64, []string) {
	n := len(sequences)
	if n == 0 {
		return 0, nil
	}

	// Create a DP table to store scores for all alignments
	dp := make([][][]float64, n)
	for i := range dp {
		dp[i] = make([][]float64, len(sequences[i])+1)
		for j := range dp[i] {
			dp[i][j] = make([]float64, len(sequences[i])+1)
		}
	}

	// Initialize DP table with gap penalties
	for i := 0; i <= len(sequences[0]); i++ {
		for j := 0; j <= len(sequences[1]); j++ {
			if i == 0 {
				dp[0][i][j] = float64(j) * gapPenalty
			} else if j == 0 {
				dp[0][i][j] = float64(i) * gapPenalty
			} else {
				match := mismatchPenalty
				if sequences[0][i-1] == sequences[1][j-1] {
					match = 0.0
				}
				dp[0][i][j] = math.Max(math.Max(dp[0][i-1][j-1]+match, dp[0][i-1][j]+gapPenalty), dp[0][i][j-1]+gapPenalty)
			}
		}
	}

	// Compute the best alignment score
	score := dp[0][len(sequences[0])][len(sequences[1])]
	aligned := make([]string, n)
	for i := range aligned {
		aligned[i] = sequences[i]
	}

	return score, aligned
}

// Print aligned sequences in a format with gaps
func printAlignment(sequences []string) {
	maxLength := 0
	for _, seq := range sequences {
		if len(seq) > maxLength {
			maxLength = len(seq)
		}
	}

	for i := range sequences {
		seq := sequences[i]
		if len(seq) < maxLength {
			seq += strings.Repeat("-", maxLength-len(seq))
		}
		sequences[i] = seq
	}

	for _, seq := range sequences {
		fmt.Println(seq)
	}
}

// Main function
func main() {
	filename := "rosalind_mult.txt"
	sequences, err := readFasta(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(sequences) < 2 {
		fmt.Println("Error: At least two sequences are required for alignment.")
		return
	}

	score, aligned := alignMultipleSequences(sequences)
	fmt.Printf("%.0f\n", score)
	printAlignment(aligned)
}
