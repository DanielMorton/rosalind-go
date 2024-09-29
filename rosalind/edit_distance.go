package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to calculate the minimum of three integers
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

// Function to compute the edit distance between two strings
func editDistance(s, t string) int {
	m := len(s)
	n := len(t)

	// Create a 2D array to store the edit distances
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize the base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] // No operation needed
			} else {
				// Take the minimum of substitution, insertion, or deletion
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	// The edit distance is in the bottom-right corner of the table
	return dp[m][n]
}

// Function to read the input file and parse the protein strings
func readFASTA(filename string) (string, string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var seqs []string
	var sb strings.Builder
	scanner := bufio.NewScanner(file)
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

	if len(seqs) != 2 {
		panic("Input file must contain exactly two sequences.")
	}

	return seqs[0], seqs[1]
}

func main() {
	// Read the protein strings from the input file
	s, t := readFASTA("rosalind_edit.txt")

	// Calculate the edit distance
	distance := editDistance(s, t)

	// Output the result
	fmt.Println(distance)
}
