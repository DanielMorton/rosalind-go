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

// immediateNeighbors generates all strings that are 1 Hamming distance away from pattern
func immediateNeighbors(pattern string) []string {
	neighbors := []string{}
	nucleotides := []byte{'A', 'C', 'G', 'T'}

	for i := 0; i < len(pattern); i++ {
		for _, nucleotide := range nucleotides {
			if nucleotide != pattern[i] {
				neighbor := pattern[:i] + string(nucleotide) + pattern[i+1:]
				neighbors = append(neighbors, neighbor)
			}
		}
	}

	return neighbors
}

// neighbors generates all strings that are at most d Hamming distance away from pattern
func neighbors(pattern string, d int) []string {
	if d == 0 {
		return []string{pattern}
	}
	if len(pattern) == 1 {
		return []string{"A", "C", "G", "T"}
	}

	neighborhood := []string{}
	suffixNeighbors := neighbors(pattern[1:], d)

	for _, suffix := range suffixNeighbors {
		if hammingDistance(pattern[1:], suffix) < d {
			for _, nucleotide := range []byte{'A', 'C', 'G', 'T'} {
				neighborhood = append(neighborhood, string(nucleotide)+suffix)
			}
		} else {
			neighborhood = append(neighborhood, string(pattern[0])+suffix)
		}
	}

	return neighborhood
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba1n.txt")
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

	// Read d
	scanner.Scan()
	d, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Error parsing d:", err)
		return
	}

	// Generate d-Neighborhood
	neighborhood := neighbors(pattern, d)

	// Print the result
	for _, neighbor := range neighborhood {
		fmt.Println(neighbor)
	}
}
