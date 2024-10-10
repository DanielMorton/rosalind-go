package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	matchScore      = 1
	mismatchPenalty = 2
	indelPenalty    = 2
)

func main() {
	// Read input from file
	v, w := readInput("rosalind_ba5i.txt")

	// Compute overlap alignment
	score, alignedV, alignedW := overlapAlignment(v, w)

	// Print results
	fmt.Println(score)
	fmt.Println(alignedV)
	fmt.Println(alignedW)
}

func readInput(filename string) (string, string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	v := scanner.Text()
	scanner.Scan()
	w := scanner.Text()

	return v, w
}

func overlapAlignment(v, w string) (int, string, string) {
	n, m := len(v), len(w)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// Initialize first row with penalties
	for j := 1; j <= m; j++ {
		dp[0][j] = -j * indelPenalty
	}

	// No need to initialize first column, it stays 0

	// Fill the dp matrix
	maxScore, maxJ := 0, 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			match := dp[i-1][j-1]
			if v[i-1] == w[j-1] {
				match += matchScore
			} else {
				match -= mismatchPenalty
			}

			dp[i][j] = max(match, dp[i-1][j]-indelPenalty, dp[i][j-1]-indelPenalty)

			// Update max score only when we reach the end of v
			if i == n && dp[i][j] >= maxScore {
				maxScore = dp[i][j]
				maxJ = j
			}
		}
	}

	// Backtrack to find the alignment
	alignedV, alignedW := "", ""
	i, j := n, maxJ
	for j > 0 {
		if i > 0 && dp[i][j] == dp[i-1][j-1]+matchScore && v[i-1] == w[j-1] {
			alignedV = string(v[i-1]) + alignedV
			alignedW = string(w[j-1]) + alignedW
			i--
			j--
		} else if i > 0 && dp[i][j] == dp[i-1][j-1]-mismatchPenalty {
			alignedV = string(v[i-1]) + alignedV
			alignedW = string(w[j-1]) + alignedW
			i--
			j--
		} else if i > 0 && dp[i][j] == dp[i-1][j]-indelPenalty {
			alignedV = string(v[i-1]) + alignedV
			alignedW = "-" + alignedW
			i--
		} else {
			alignedV = "-" + alignedV
			alignedW = string(w[j-1]) + alignedW
			j--
		}
	}

	return maxScore, alignedV, alignedW
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}
