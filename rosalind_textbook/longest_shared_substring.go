package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func longestSharedSubstring(text1, text2 string) string {
	m, n := len(text1), len(text2)
	longest := ""

	// Create a 2D slice to store the lengths of common substrings
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > len(longest) {
					longest = text1[i-dp[i][j] : i]
				}
			}
		}
	}

	return longest
}

func main() {
	// Read input from file
	content, err := ioutil.ReadFile("rosalind_ba9e.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the content into two strings
	texts := strings.Split(strings.TrimSpace(string(content)), "\n")
	if len(texts) != 2 {
		fmt.Println("Invalid input format")
		return
	}

	text1, text2 := texts[0], texts[1]

	// Find the longest shared substring
	result := longestSharedSubstring(text1, text2)

	// Print the result
	fmt.Println(result)
}
