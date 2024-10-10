package main

import (
	"bufio"
	"fmt"
	"os"
)

func longestCommonSubsequence(s1, s2 string) string {
	m, n := len(s1), len(s2)

	// Create a 2D slice to store the lengths of LCS
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Backtrack to find the LCS
	lcs := ""
	i, j := m, n
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			lcs = string(s1[i-1]) + lcs
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return lcs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Read input from file
	file, err := os.Open("rosalind_ba5c.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the two strings
	var s1, s2 string
	if scanner.Scan() {
		s1 = scanner.Text()
	}
	if scanner.Scan() {
		s2 = scanner.Text()
	}

	// Calculate and print the result
	result := longestCommonSubsequence(s1, s2)
	fmt.Println(result)
}
