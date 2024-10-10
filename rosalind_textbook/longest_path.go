package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestPath(n, m int, down [][]int, right [][]int) int {
	// Create a 2D slice to store the longest path lengths
	s := make([][]int, n+1)
	for i := range s {
		s[i] = make([]int, m+1)
	}

	// Initialize the first row and column
	for i := 1; i <= n; i++ {
		s[i][0] = s[i-1][0] + down[i-1][0]
	}
	for j := 1; j <= m; j++ {
		s[0][j] = s[0][j-1] + right[0][j-1]
	}

	// Fill the rest of the matrix
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			s[i][j] = max(s[i-1][j]+down[i-1][j], s[i][j-1]+right[i][j-1])
		}
	}

	// Return the length of the longest path to the sink
	return s[n][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Read input from file
	file, err := os.Open("rosalind_ba5b.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read n and m
	scanner.Scan()
	dims := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(dims[0])
	m, _ := strconv.Atoi(dims[1])

	// Read Down matrix
	down := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		row := strings.Fields(scanner.Text())
		down[i] = make([]int, m+1)
		for j, val := range row {
			down[i][j], _ = strconv.Atoi(val)
		}
	}

	// Skip the "-" line
	scanner.Scan()

	// Read Right matrix
	right := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		scanner.Scan()
		row := strings.Fields(scanner.Text())
		right[i] = make([]int, m)
		for j, val := range row {
			right[i][j], _ = strconv.Atoi(val)
		}
	}

	// Calculate and print the result
	result := longestPath(n, m, down, right)
	fmt.Println(result)
}
