package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	match    = 1
	mismatch = -2
	gap      = -2
)

func main() {
	// Read input from file
	s, t := readFASTA("rosalind_oap.txt")

	// Perform overlap alignment
	score, aligned1, aligned2 := overlapAlignment(s, t)

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
	var currentSeq *string

	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '>' {
			if s == "" {
				currentSeq = &s
			} else {
				currentSeq = &t
			}
		} else {
			*currentSeq += line
		}
	}

	return s, t
}

func overlapAlignment(s, t string) (int, string, string) {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize first column (free gaps at the beginning of s)
	for i := 0; i <= m; i++ {
		dp[i][0] = 0
	}

	// Fill the dp table
	maxScore, maxI, maxJ := 0, 0, 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			match := dp[i-1][j-1]
			if s[i-1] == t[j-1] {
				match += 1
			} else {
				match -= 2
			}
			dp[i][j] = max(match, dp[i-1][j]+gap, dp[i][j-1]+gap)

			// Update max score for each position in the last row
			if i == m && dp[i][j] >= maxScore {
				maxScore = dp[i][j]
				maxI, maxJ = i, j
			}
		}
	}

	// Traceback
	i, j := maxI, maxJ
	var aligned1, aligned2 strings.Builder

	for i > 0 && j > 0 {
		if j > 0 && dp[i][j] == dp[i][j-1]+gap {
			aligned1.WriteByte('-')
			aligned2.WriteByte(t[j-1])
			j--
		} else if i > 0 && dp[i][j] == dp[i-1][j]+gap {
			aligned1.WriteByte(s[i-1])
			aligned2.WriteByte('-')
			i--
		} else {
			aligned1.WriteByte(s[i-1])
			aligned2.WriteByte(t[j-1])
			i--
			j--
		}
	}

	// Add any remaining characters from the prefix of t
	for j > 0 {
		aligned1.WriteByte('-')
		aligned2.WriteByte(t[j-1])
		j--
	}

	return maxScore, reverse(aligned1.String()), reverse(aligned2.String())
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
