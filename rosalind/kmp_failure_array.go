package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// computeFailureArray computes the failure array P for the given string s
func computeFailureArray(s string) []int {
	n := len(s)
	P := make([]int, n)

	// Initialize the first element
	P[0] = 0

	// Compute the failure array
	k := 0
	for i := 1; i < n; i++ {
		for k > 0 && s[k] != s[i] {
			k = P[k-1]
		}
		if s[k] == s[i] {
			k++
		}
		P[i] = k
	}

	return P
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_kmp.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sequence strings.Builder

	// Read the sequence from the FASTA file
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, ">") { // Skip header lines
			sequence.WriteString(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Compute the failure array
	s := sequence.String()
	failureArray := computeFailureArray(s)

	// Print the failure array
	for i, value := range failureArray {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(value)
	}
	fmt.Println()
}
