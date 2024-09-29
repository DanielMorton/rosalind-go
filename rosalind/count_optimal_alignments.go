package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const MOD = 134217727

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func countAlignments(s, t string) int {
	m, n := len(s), len(t)

	// Initialize dp and count tables
	dp := make([][]int, m+1)
	count := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		count[i] = make([]int, n+1)
	}

	// Base case initialization
	for i := 0; i <= m; i++ {
		dp[i][0] = i
		count[i][0] = 1
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
		count[0][j] = 1
	}

	// Fill dp and count tables
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cost := 1
			if s[i-1] == t[j-1] {
				cost = 0
			}
			dp[i][j] = min(dp[i-1][j-1]+cost, dp[i-1][j]+1, dp[i][j-1]+1)

			count[i][j] = 0
			if dp[i][j] == dp[i-1][j-1]+cost {
				count[i][j] += count[i-1][j-1]
			}
			if dp[i][j] == dp[i-1][j]+1 {
				count[i][j] += count[i-1][j]
			}
			if dp[i][j] == dp[i][j-1]+1 {
				count[i][j] += count[i][j-1]
			}
			count[i][j] %= MOD
		}
	}

	return count[m][n]
}

func main() {
	file, err := os.Open("rosalind_ctea.txt")
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

	result := countAlignments(seq1, seq2)

	fmt.Println(result)
}
