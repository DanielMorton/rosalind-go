package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	MATCH_SCORE      = 1
	MISMATCH_PENALTY = -1
	GAP_PENALTY      = -1
)

func fittingAlignment(v, w string) (int, string, string) {
	m, n := len(v), len(w)

	// Create a 2D slice to store the alignment scores
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize the first row with zeros (allowing free starting point in v)
	for j := 1; j <= n; j++ {
		dp[0][j] = j * GAP_PENALTY
	}

	// Fill the dp table
	maxScore, maxIndex := 0, 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			match := dp[i-1][j-1]
			if v[i-1] == w[j-1] {
				match += MATCH_SCORE
			} else {
				match += MISMATCH_PENALTY
			}

			dp[i][j] = max(match, dp[i-1][j]+GAP_PENALTY, dp[i][j-1]+GAP_PENALTY)

			// Update max score if we've aligned all of w
			if j == n && dp[i][j] > maxScore {
				maxScore = dp[i][j]
				maxIndex = i
			}
		}
	}

	// Traceback
	alignV, alignW := "", ""
	i, j := maxIndex, n
	for j > 0 {
		if i > 0 && dp[i][j] == dp[i-1][j-1]+scoreFunction(v[i-1], w[j-1]) {
			alignV = string(v[i-1]) + alignV
			alignW = string(w[j-1]) + alignW
			i--
			j--
		} else if i > 0 && dp[i][j] == dp[i-1][j]+GAP_PENALTY {
			alignV = string(v[i-1]) + alignV
			alignW = "-" + alignW
			i--
		} else {
			alignV = "-" + alignV
			alignW = string(w[j-1]) + alignW
			j--
		}
	}

	return maxScore, alignV, alignW
}

func scoreFunction(a, b byte) int {
	if a == b {
		return MATCH_SCORE
	}
	return MISMATCH_PENALTY
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}

func main() {
	// Read input from file
	file, err := os.Open("rosalind_ba5h.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var v, w string
	if scanner.Scan() {
		v = strings.TrimSpace(scanner.Text())
	}
	if scanner.Scan() {
		w = strings.TrimSpace(scanner.Text())
	}

	// Calculate and print the fitting alignment
	score, alignV, alignW := fittingAlignment(v, w)
	fmt.Println(score)
	fmt.Println(alignV)
	fmt.Println(alignW)
}
