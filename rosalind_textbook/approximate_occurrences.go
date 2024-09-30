package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// hammingDistance calculates the Hamming distance between two strings
func hammingDistance(s1, s2 string) int {
	distance := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			distance++
		}
	}
	return distance
}

// approximatePatternMatch finds all starting positions where Pattern appears in Text with at most d mismatches
func approximatePatternMatch(pattern, text string, d int) []int {
	var positions []int
	k := len(pattern)

	for i := 0; i <= len(text)-k; i++ {
		substring := text[i : i+k]
		if hammingDistance(pattern, substring) <= d {
			positions = append(positions, i)
		}
	}

	return positions
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba1h.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read input from file
	scanner := bufio.NewScanner(file)

	// Read Pattern
	scanner.Scan()
	pattern := strings.TrimSpace(scanner.Text())

	// Read Text
	scanner.Scan()
	text := strings.TrimSpace(scanner.Text())

	// Read d
	scanner.Scan()
	d, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Error parsing d:", err)
		return
	}

	// Find approximate pattern matches
	positions := approximatePatternMatch(pattern, text, d)

	// Print the result
	for i, pos := range positions {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(pos)
	}
	fmt.Println()
}
