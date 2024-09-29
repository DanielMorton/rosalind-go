package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var blosum62 = map[string]map[string]int{
	"A": {"A": 4, "C": 0, "D": -2, "E": -1, "F": -2, "G": 0, "H": -2, "I": -1, "K": -1, "L": -1, "M": -1, "N": -2, "P": -1, "Q": -1, "R": -1, "S": 1, "T": 0, "V": 0, "W": -3, "Y": -2},
	"C": {"A": 0, "C": 9, "D": -3, "E": -4, "F": -2, "G": -3, "H": -3, "I": -1, "K": -3, "L": -1, "M": -1, "N": -3, "P": -3, "Q": -3, "R": -3, "S": -1, "T": -1, "V": -1, "W": -2, "Y": -2},
	"D": {"A": -2, "C": -3, "D": 6, "E": 2, "F": -3, "G": -1, "H": -1, "I": -3, "K": -1, "L": -4, "M": -3, "N": 1, "P": -1, "Q": 0, "R": -2, "S": 0, "T": -1, "V": -3, "W": -4, "Y": -3},
	"E": {"A": -1, "C": -4, "D": 2, "E": 5, "F": -3, "G": -2, "H": 0, "I": -3, "K": 1, "L": -3, "M": -2, "N": 0, "P": -1, "Q": 2, "R": 0, "S": 0, "T": -1, "V": -2, "W": -3, "Y": -2},
	"F": {"A": -2, "C": -2, "D": -3, "E": -3, "F": 6, "G": -3, "H": -1, "I": 0, "K": -3, "L": 0, "M": 0, "N": -3, "P": -4, "Q": -3, "R": -3, "S": -2, "T": -2, "V": -1, "W": 1, "Y": 3},
	"G": {"A": 0, "C": -3, "D": -1, "E": -2, "F": -3, "G": 6, "H": -2, "I": -4, "K": -2, "L": -4, "M": -3, "N": 0, "P": -2, "Q": -2, "R": -2, "S": 0, "T": -2, "V": -3, "W": -2, "Y": -3},
	"H": {"A": -2, "C": -3, "D": -1, "E": 0, "F": -1, "G": -2, "H": 8, "I": -3, "K": -1, "L": -3, "M": -2, "N": 1, "P": -2, "Q": 0, "R": 0, "S": -1, "T": -2, "V": -3, "W": -2, "Y": 2},
	"I": {"A": -1, "C": -1, "D": -3, "E": -3, "F": 0, "G": -4, "H": -3, "I": 4, "K": -3, "L": 2, "M": 1, "N": -3, "P": -3, "Q": -3, "R": -3, "S": -2, "T": -1, "V": 3, "W": -3, "Y": -1},
	"K": {"A": -1, "C": -3, "D": -1, "E": 1, "F": -3, "G": -2, "H": -1, "I": -3, "K": 5, "L": -2, "M": -1, "N": 0, "P": -1, "Q": 1, "R": 2, "S": 0, "T": -1, "V": -2, "W": -3, "Y": -2},
	"L": {"A": -1, "C": -1, "D": -4, "E": -3, "F": 0, "G": -4, "H": -3, "I": 2, "K": -2, "L": 4, "M": 2, "N": -3, "P": -3, "Q": -2, "R": -2, "S": -2, "T": -1, "V": 1, "W": -2, "Y": -1},
	"M": {"A": -1, "C": -1, "D": -3, "E": -2, "F": 0, "G": -3, "H": -2, "I": 1, "K": -1, "L": 2, "M": 5, "N": -2, "P": -2, "Q": 0, "R": -1, "S": -1, "T": -1, "V": 1, "W": -1, "Y": -1},
	"N": {"A": -2, "C": -3, "D": 1, "E": 0, "F": -3, "G": 0, "H": 1, "I": -3, "K": 0, "L": -3, "M": -2, "N": 6, "P": -2, "Q": 0, "R": 0, "S": 1, "T": 0, "V": -3, "W": -4, "Y": -2},
	"P": {"A": -1, "C": -3, "D": -1, "E": -1, "F": -4, "G": -2, "H": -2, "I": -3, "K": -1, "L": -3, "M": -2, "N": -2, "P": 7, "Q": -1, "R": -2, "S": -1, "T": -1, "V": -2, "W": -4, "Y": -3},
	"Q": {"A": -1, "C": -3, "D": 0, "E": 2, "F": -3, "G": -2, "H": 0, "I": -3, "K": 1, "L": -2, "M": 0, "N": 0, "P": -1, "Q": 5, "R": 1, "S": 0, "T": -1, "V": -2, "W": -2, "Y": -1},
	"R": {"A": -1, "C": -3, "D": -2, "E": 0, "F": -3, "G": -2, "H": 0, "I": -3, "K": 2, "L": -2, "M": -1, "N": 0, "P": -2, "Q": 1, "R": 5, "S": -1, "T": -1, "V": -3, "W": -3, "Y": -2},
	"S": {"A": 1, "C": -1, "D": 0, "E": 0, "F": -2, "G": 0, "H": -1, "I": -2, "K": 0, "L": -2, "M": -1, "N": 1, "P": -1, "Q": 0, "R": -1, "S": 4, "T": 1, "V": -2, "W": -3, "Y": -2},
	"T": {"A": 0, "C": -1, "D": -1, "E": -1, "F": -2, "G": -2, "H": -2, "I": -1, "K": -1, "L": -1, "M": -1, "N": 0, "P": -1, "Q": -1, "R": -1, "S": 1, "T": 5, "V": 0, "W": -2, "Y": -2},
	"V": {"A": 0, "C": -1, "D": -3, "E": -2, "F": -1, "G": -3, "H": -3, "I": 3, "K": -2, "L": 1, "M": 1, "N": -3, "P": -2, "Q": -2, "R": -3, "S": -2, "T": 0, "V": 4, "W": -3, "Y": -1},
	"W": {"A": -3, "C": -2, "D": -4, "E": -3, "F": 1, "G": -2, "H": -2, "I": -3, "K": -3, "L": -2, "M": -1, "N": -4, "P": -4, "Q": -2, "R": -3, "S": -3, "T": -2, "V": -3, "W": 11, "Y": 2},
	"Y": {"A": -2, "C": -2, "D": -3, "E": -2, "F": 3, "G": -3, "H": 2, "I": -1, "K": -2, "L": -1, "M": -1, "N": -2, "P": -3, "Q": -1, "R": -2, "S": -2, "T": -2, "V": -1, "W": 2, "Y": 7},
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func needlemanWunsch(s1, s2 string, gapPenalty int) (int, string, string) {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize the dp array for gaps
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + gapPenalty
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] + gapPenalty
	}

	// Fill the dp array with BLOSUM62 scores
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			match := dp[i-1][j-1] + blosum62[string(s1[i-1])][string(s2[j-1])]
			delete := dp[i-1][j] + gapPenalty
			insert := dp[i][j-1] + gapPenalty
			dp[i][j] = max(match, delete, insert)
		}
	}

	// Backtrack to get the aligned sequences
	alignedS1, alignedS2 := "", ""
	i, j := m, n
	for i > 0 || j > 0 {
		if i > 0 && j > 0 && dp[i][j] == dp[i-1][j-1]+blosum62[string(s1[i-1])][string(s2[j-1])] {
			alignedS1 = string(s1[i-1]) + alignedS1
			alignedS2 = string(s2[j-1]) + alignedS2
			i--
			j--
		} else if i > 0 && dp[i][j] == dp[i-1][j]+gapPenalty {
			alignedS1 = string(s1[i-1]) + alignedS1
			alignedS2 = "-" + alignedS2
			i--
		} else {
			alignedS1 = "-" + alignedS1
			alignedS2 = string(s2[j-1]) + alignedS2
			j--
		}
	}

	return dp[m][n], alignedS1, alignedS2
}

func main() {
	file, err := os.Open("rosalind_glob.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sequences := []string{}
	currentSeq := ""

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || line[0] == '>' {
			if currentSeq != "" {
				sequences = append(sequences, currentSeq)
				currentSeq = ""
			}
		} else {
			currentSeq += line
		}
	}

	if currentSeq != "" {
		sequences = append(sequences, currentSeq)
	}

	if len(sequences) != 2 {
		log.Fatalf("Expected 2 sequences, but got %d", len(sequences))
	}

	seq1, seq2 := sequences[0], sequences[1]
	gapPenalty := -5

	score, _, _ := needlemanWunsch(seq1, seq2, gapPenalty)

	fmt.Println(score)
}
