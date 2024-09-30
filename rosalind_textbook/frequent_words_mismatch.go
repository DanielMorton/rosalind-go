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

// generateAllKmers generates all possible k-mers
func generateAllKmers(k int) []string {
	nucleotides := []byte{'A', 'C', 'G', 'T'}
	kmers := []string{""}
	for i := 0; i < k; i++ {
		var newKmers []string
		for _, kmer := range kmers {
			for _, nucleotide := range nucleotides {
				newKmers = append(newKmers, kmer+string(nucleotide))
			}
		}
		kmers = newKmers
	}
	return kmers
}

// frequentWordsWithMismatches finds the most frequent k-mers with up to d mismatches in Text
func frequentWordsWithMismatches(text string, k, d int) []string {
	kmerCounts := make(map[string]int)
	allKmers := generateAllKmers(k)

	// Count occurrences of each k-mer with up to d mismatches
	for _, kmer := range allKmers {
		for i := 0; i <= len(text)-k; i++ {
			if hammingDistance(kmer, text[i:i+k]) <= d {
				kmerCounts[kmer]++
			}
		}
	}

	// Find the maximum count
	maxCount := 0
	for _, count := range kmerCounts {
		if count > maxCount {
			maxCount = count
		}
	}

	// Collect all k-mers with the maximum count
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
	file, err := os.Open("rosalind_ba1i.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read input from file
	scanner := bufio.NewScanner(file)

	// Read Text
	scanner.Scan()
	text := strings.TrimSpace(scanner.Text())

	// Read k and d
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	if len(params) != 2 {
		fmt.Println("Invalid input format")
		return
	}
	k, _ := strconv.Atoi(params[0])
	d, _ := strconv.Atoi(params[1])

	// Find frequent words with mismatches
	result := frequentWordsWithMismatches(text, k, d)

	// Print the result
	fmt.Println(strings.Join(result, " "))
}
