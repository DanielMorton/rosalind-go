package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func editDistance(s1, s2 string) int {
	m, n := len(s1), len(s2)

	// Create a 2D slice to store the edit distances
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize the first row and column
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] // No operation needed
			} else {
				dp[i][j] = min(
					dp[i-1][j]+1,   // Deletion
					dp[i][j-1]+1,   // Insertion
					dp[i-1][j-1]+1, // Substitution
				)
			}
		}
	}

	// The bottom-right cell contains the edit distance
	return dp[m][n]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}

func main() {
	// Read input from file
	file, err := os.Open("rosalind_ba5g.txt")
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

	// Calculate and print the edit distance
	distance := editDistance(s1, s2)
	fmt.Println(distance)
}
