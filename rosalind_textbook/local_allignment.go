package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PAM250 scoring matrix
var PAM250 = map[string]map[string]int{
	"A": {"A": 2, "R": -2, "N": 0, "D": 0, "C": -2, "Q": 0, "E": 0, "G": 1, "H": -1, "I": -1, "L": -2, "K": -1, "M": -1, "F": -3, "P": 1, "S": 1, "T": 1, "W": -6, "Y": -3, "V": 0},
	"R": {"A": -2, "R": 6, "N": 0, "D": -1, "C": -4, "Q": 1, "E": -1, "G": -3, "H": 2, "I": -2, "L": -3, "K": 3, "M": 0, "F": -4, "P": 0, "S": 0, "T": -1, "W": 2, "Y": -4, "V": -2},
	"N": {"A": 0, "R": 0, "N": 2, "D": 2, "C": -4, "Q": 1, "E": 1, "G": 0, "H": 2, "I": -2, "L": -3, "K": 1, "M": -2, "F": -3, "P": 0, "S": 1, "T": 0, "W": -4, "Y": -2, "V": -2},
	"D": {"A": 0, "R": -1, "N": 2, "D": 4, "C": -5, "Q": 2, "E": 3, "G": 1, "H": 1, "I": -2, "L": -4, "K": 0, "M": -3, "F": -6, "P": -1, "S": 0, "T": 0, "W": -7, "Y": -4, "V": -2},
	"C": {"A": -2, "R": -4, "N": -4, "D": -5, "C": 12, "Q": -5, "E": -5, "G": -3, "H": -3, "I": -2, "L": -6, "K": -5, "M": -5, "F": -4, "P": -3, "S": 0, "T": -2, "W": -8, "Y": 0, "V": -2},
	"Q": {"A": 0, "R": 1, "N": 1, "D": 2, "C": -5, "Q": 4, "E": 2, "G": -1, "H": 3, "I": -2, "L": -2, "K": 1, "M": -1, "F": -5, "P": 0, "S": -1, "T": -1, "W": -5, "Y": -4, "V": -2},
	"E": {"A": 0, "R": -1, "N": 1, "D": 3, "C": -5, "Q": 2, "E": 4, "G": 0, "H": 1, "I": -2, "L": -3, "K": 0, "M": -2, "F": -5, "P": -1, "S": 0, "T": 0, "W": -7, "Y": -4, "V": -2},
	"G": {"A": 1, "R": -3, "N": 0, "D": 1, "C": -3, "Q": -1, "E": 0, "G": 5, "H": -2, "I": -3, "L": -4, "K": -2, "M": -3, "F": -5, "P": 0, "S": 1, "T": 0, "W": -7, "Y": -5, "V": -1},
	"H": {"A": -1, "R": 2, "N": 2, "D": 1, "C": -3, "Q": 3, "E": 1, "G": -2, "H": 6, "I": -2, "L": -2, "K": 0, "M": -2, "F": -2, "P": 0, "S": -1, "T": -1, "W": -3, "Y": 0, "V": -2},
	"I": {"A": -1, "R": -2, "N": -2, "D": -2, "C": -2, "Q": -2, "E": -2, "G": -3, "H": -2, "I": 5, "L": 2, "K": -2, "M": 2, "F": 1, "P": -2, "S": -1, "T": 0, "W": -5, "Y": -1, "V": 4},
	"L": {"A": -2, "R": -3, "N": -3, "D": -4, "C": -6, "Q": -2, "E": -3, "G": -4, "H": -2, "I": 2, "L": 6, "K": -3, "M": 4, "F": 2, "P": -3, "S": -3, "T": -2, "W": -2, "Y": -1, "V": 2},
	"K": {"A": -1, "R": 3, "N": 1, "D": 0, "C": -5, "Q": 1, "E": 0, "G": -2, "H": 0, "I": -2, "L": -3, "K": 5, "M": 0, "F": -5, "P": -1, "S": 0, "T": 0, "W": -3, "Y": -4, "V": -2},
	"M": {"A": -1, "R": 0, "N": -2, "D": -3, "C": -5, "Q": -1, "E": -2, "G": -3, "H": -2, "I": 2, "L": 4, "K": 0, "M": 6, "F": 0, "P": -2, "S": -2, "T": -1, "W": -4, "Y": -2, "V": 2},
	"F": {"A": -3, "R": -4, "N": -3, "D": -6, "C": -4, "Q": -5, "E": -5, "G": -5, "H": -2, "I": 1, "L": 2, "K": -5, "M": 0, "F": 9, "P": -5, "S": -3, "T": -3, "W": 0, "Y": 7, "V": -1},
	"P": {"A": 1, "R": 0, "N": 0, "D": -1, "C": -3, "Q": 0, "E": -1, "G": 0, "H": 0, "I": -2, "L": -3, "K": -1, "M": -2, "F": -5, "P": 6, "S": 1, "T": 0, "W": -6, "Y": -5, "V": -1},
	"S": {"A": 1, "R": 0, "N": 1, "D": 0, "C": 0, "Q": -1, "E": 0, "G": 1, "H": -1, "I": -1, "L": -3, "K": 0, "M": -2, "F": -3, "P": 1, "S": 2, "T": 1, "W": -2, "Y": -3, "V": -1},
	"T": {"A": 1, "R": -1, "N": 0, "D": 0, "C": -2, "Q": -1, "E": 0, "G": 0, "H": -1, "I": 0, "L": -2, "K": 0, "M": -1, "F": -3, "P": 0, "S": 1, "T": 3, "W": -5, "Y": -3, "V": 0},
	"W": {"A": -6, "R": 2, "N": -4, "D": -7, "C": -8, "Q": -5, "E": -7, "G": -7, "H": -3, "I": -5, "L": -2, "K": -3, "M": -4, "F": 0, "P": -6, "S": -2, "T": -5, "W": 17, "Y": 0, "V": -6},
	"Y": {"A": -3, "R": -4, "N": -2, "D": -4, "C": 0, "Q": -4, "E": -4, "G": -5, "H": 0, "I": -1, "L": -1, "K": -4, "M": -2, "F": 7, "P": -5, "S": -3, "T": -3, "W": 0, "Y": 10, "V": -2},
	"V": {"A": 0, "R": -2, "N": -2, "D": -2, "C": -2, "Q": -2, "E": -2, "G": -1, "H": -2, "I": 4, "L": 2, "K": -2, "M": 2, "F": -1, "P": -1, "S": -1, "T": 0, "W": -6, "Y": -2, "V": 4},
}

const GAP_PENALTY = -5

func localAlignment(s1, s2 string) (int, string, string) {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Variables to keep track of the highest score and its position
	maxScore, maxI, maxJ := 0, 0, 0

	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			match := dp[i-1][j-1] + PAM250[string(s1[i-1])][string(s2[j-1])]
			delete := dp[i-1][j] + GAP_PENALTY
			insert := dp[i][j-1] + GAP_PENALTY
			dp[i][j] = max(0, max(match, max(delete, insert)))

			// Update max score if necessary
			if dp[i][j] > maxScore {
				maxScore = dp[i][j]
				maxI, maxJ = i, j
			}
		}
	}

	// Traceback
	align1, align2 := "", ""
	i, j := maxI, maxJ
	for i > 0 && j > 0 && dp[i][j] > 0 {
		if dp[i][j] == dp[i-1][j-1]+PAM250[string(s1[i-1])][string(s2[j-1])] {
			align1 = string(s1[i-1]) + align1
			align2 = string(s2[j-1]) + align2
			i--
			j--
		} else if dp[i][j] == dp[i-1][j]+GAP_PENALTY {
			align1 = string(s1[i-1]) + align1
			align2 = "-" + align2
			i--
		} else {
			align1 = "-" + align1
			align2 = string(s2[j-1]) + align2
			j--
		}
	}

	return maxScore, align1, align2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Read input from file
	file, err := os.Open("rosalind_ba5f.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var s1, s2 string
	if scanner.Scan() {
		s1 = strings.TrimSpace(scanner.Text())
	}
	if scanner.Scan() {
		s2 = strings.TrimSpace(scanner.Text())
	}

	score, align1, align2 := localAlignment(s1, s2)

	fmt.Println(score)
	fmt.Println(align1)
	fmt.Println(align2)
}
