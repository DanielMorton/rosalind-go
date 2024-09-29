package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readFASTA reads a FASTA file and returns the two sequences in it.
func readFASTA(filename string) (string, string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var seqs []string
	var sb strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '>' {
			if sb.Len() > 0 {
				seqs = append(seqs, sb.String())
				sb.Reset()
			}
		} else {
			sb.WriteString(line)
		}
	}
	if sb.Len() > 0 {
		seqs = append(seqs, sb.String())
	}

	return seqs[0], seqs[1]
}

// min returns the minimum of three integers.
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// computeEditDistance builds the DP table and backtracks to return the edit distance and aligned strings.
func computeEditDistance(s, t string) (int, string, string) {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize DP table
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	// Fill DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}

	// Backtrack to find aligned strings
	var alignedS, alignedT strings.Builder
	i, j := m, n
	for i > 0 || j > 0 {
		if i > 0 && j > 0 && s[i-1] == t[j-1] {
			alignedS.WriteByte(s[i-1])
			alignedT.WriteByte(t[j-1])
			i--
			j--
		} else if i > 0 && (j == 0 || dp[i][j] == dp[i-1][j]+1) {
			alignedS.WriteByte(s[i-1])
			alignedT.WriteByte('-')
			i--
		} else if j > 0 && (i == 0 || dp[i][j] == dp[i][j-1]+1) {
			alignedS.WriteByte('-')
			alignedT.WriteByte(t[j-1])
			j--
		} else {
			alignedS.WriteByte(s[i-1])
			alignedT.WriteByte(t[j-1])
			i--
			j--
		}
	}

	// Reverse the aligned strings
	alignedSStr := reverseString(alignedS.String())
	alignedTStr := reverseString(alignedT.String())

	return dp[m][n], alignedSStr, alignedTStr
}

// reverseString reverses a string.
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	s, t := readFASTA("rosalind_edta.txt")
	distance, alignedS, alignedT := computeEditDistance(s, t)
	fmt.Println(distance)
	fmt.Println(alignedS)
	fmt.Println(alignedT)
}
