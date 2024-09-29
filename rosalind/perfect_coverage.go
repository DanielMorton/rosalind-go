package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read k-mers from file
	kmers := readKmers("rosalind_pcov.txt")

	// Find the cyclic superstring
	superstring := findCyclicSuperstring(kmers)

	fmt.Println(superstring)
}

func readKmers(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var kmers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		kmer := strings.TrimSpace(scanner.Text())
		if kmer != "" {
			kmers = append(kmers, kmer)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return kmers
}

func findCyclicSuperstring(kmers []string) string {
	// Create a map to store the overlaps
	overlapMap := make(map[string]string)
	for _, kmer := range kmers {
		prefix := kmer[:len(kmer)-1]
		overlapMap[prefix] = kmer[len(kmer)-1:]
	}

	// Start with any k-mer
	start := kmers[0][:len(kmers[0])-1]
	result := start

	for i := 0; i < len(kmers); i++ {
		result += overlapMap[start]
		start = result[len(result)-len(kmers[0])+1:]
	}

	// Trim the result to get the minimal cyclic superstring
	return result[:len(kmers)]
}
