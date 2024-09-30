package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMostFrequentKmers(text string, k int) []string {
	kmerCounts := make(map[string]int)
	maxCount := 0

	// Count occurrences of each k-mer
	for i := 0; i <= len(text)-k; i++ {
		kmer := text[i : i+k]
		kmerCounts[kmer]++
		if kmerCounts[kmer] > maxCount {
			maxCount = kmerCounts[kmer]
		}
	}

	// Find all k-mers with the maximum count
	var mostFrequent []string
	for kmer, count := range kmerCounts {
		if count == maxCount {
			mostFrequent = append(mostFrequent, kmer)
		}
	}

	return mostFrequent
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba1b.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read input from file
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := strings.TrimSpace(scanner.Text())
	scanner.Scan()
	k, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Error parsing k:", err)
		return
	}

	// Find the most frequent k-mers
	result := findMostFrequentKmers(text, k)

	// Print the result
	fmt.Println(strings.Join(result, " "))
}
