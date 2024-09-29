package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("rosalind_lcsq.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read input file
	scanner := bufio.NewScanner(file)
	var id, s, t string
	var sequences []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if id != "" {
				sequences = append(sequences, s)
			}
			id = line
			s = ""
		} else {
			s += line
		}
	}
	sequences = append(sequences, s) // Append the last sequence

	if len(sequences) != 2 {
		log.Fatal("The input file should contain exactly two DNA sequences.")
	}
	s, t = sequences[0], sequences[1]

	// Compute the longest common subsequence
	lcs := longestCommonSubsequence(s, t)
	fmt.Println(lcs)
}

func longestCommonSubsequence(s, t string) string {
	m, n := len(s), len(t)
	dp := make([][]string, m+1)
	for i := range dp {
		dp[i] = make([]string, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + string(s[i-1])
			} else {
				if len(dp[i-1][j]) > len(dp[i][j-1]) {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[m][n]
}
