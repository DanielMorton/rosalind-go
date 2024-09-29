package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function to find the longest common subsequence (LCS)
func LCS(s, t string) [][]int {
	m, n := len(s), len(t)
	// Create a 2D array to store lengths of LCS for substrings
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill dp array
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp
}

// Function to reconstruct the shortest common supersequence
func SCS(s, t string) string {
	dp := LCS(s, t)
	m, n := len(s), len(t)
	i, j := m, n
	var scs []byte

	// Trace back through the dp array to build the supersequence
	for i > 0 && j > 0 {
		if s[i-1] == t[j-1] {
			scs = append([]byte{s[i-1]}, scs...)
			i--
			j--
		} else if dp[i-1][j] >= dp[i][j-1] {
			scs = append([]byte{s[i-1]}, scs...)
			i--
		} else {
			scs = append([]byte{t[j-1]}, scs...)
			j--
		}
	}

	// Add remaining characters from either string
	for i > 0 {
		scs = append([]byte{s[i-1]}, scs...)
		i--
	}
	for j > 0 {
		scs = append([]byte{t[j-1]}, scs...)
		j--
	}

	return string(scs)
}

// Helper function to find max of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Read the input file
	file, err := os.Open("rosalind_scsp.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var s, t string
	if scanner.Scan() {
		s = scanner.Text()
	}
	if scanner.Scan() {
		t = scanner.Text()
	}

	// Calculate and print the shortest common supersequence
	result := SCS(s, t)
	fmt.Println(result)
}
