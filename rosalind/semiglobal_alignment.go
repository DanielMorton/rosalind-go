package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	match    = 1
	mismatch = -1
	gap      = -1
)

func main() {
	// Read input from file
	s, t := readFASTA("rosalind_smgb.txt")

	// Perform semiglobal alignment
	score, aligned1, aligned2 := semiglobalAlign(s, t)

	// Print results
	fmt.Println(score)
	fmt.Println(aligned1)
	fmt.Println(aligned2)
}

func readFASTA(filename string) (string, string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var s, t string
	currentSeq := ""

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if currentSeq != "" {
				s = currentSeq
				currentSeq = ""
			}
		} else {
			currentSeq += line
		}
	}
	t = currentSeq

	return s, t
}

func semiglobalAlign(s, t string) (int, string, string) {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			match := dp[i-1][j-1]
			if s[i-1] == t[j-1] {
				match += 1
			} else {
				match += mismatch
			}

			// Use 0 for gap penalty if we're at the end of either sequence
			gapS := dp[i-1][j]
			if j == n {
				gapS = dp[i-1][j]
			} else {
				gapS = dp[i-1][j] + gap
			}

			gapT := dp[i][j-1]
			if i == m {
				gapT = dp[i][j-1]
			} else {
				gapT = dp[i][j-1] + gap
			}

			dp[i][j] = max(match, gapS, gapT)
		}
	}

	// Find the maximum score
	maxScore := dp[m][n]
	endI, endJ := m, n

	// Traceback
	aligned1, aligned2 := "", ""
	i, j := endI, endJ
	for i > 0 && j > 0 {
		diag := dp[i-1][j-1]
		if s[i-1] == t[j-1] {
			diag += match
		} else {
			diag += mismatch
		}

		if dp[i][j] == diag {
			aligned1 = string(s[i-1]) + aligned1
			aligned2 = string(t[j-1]) + aligned2
			i--
			j--
		} else if dp[i][j] == dp[i][j-1] && i == m ||
			(dp[i][j] == dp[i][j-1]+gap && i < m) {
			aligned1 = "-" + aligned1
			aligned2 = string(t[j-1]) + aligned2
			j--
		} else if dp[i][j] == dp[i-1][j] && j == n ||
			(dp[i][j] == dp[i-1][j]+gap && j < n) {
			aligned1 = string(s[i-1]) + aligned1
			aligned2 = "-" + aligned2
			i--
		}
	}

	// Add leading gaps (free in semiglobal alignment)
	for i > 0 {
		aligned1 = string(s[i-1]) + aligned1
		aligned2 = "-" + aligned2
		i--
	}
	for j > 0 {
		aligned1 = "-" + aligned1
		aligned2 = string(t[j-1]) + aligned2
		j--
	}

	return maxScore, aligned1, aligned2
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}
