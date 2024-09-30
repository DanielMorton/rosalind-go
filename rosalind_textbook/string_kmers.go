package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func generateKmers(text string, k int) []string {
	kmers := make([]string, 0)
	for i := 0; i <= len(text)-k; i++ {
		kmer := text[i : i+k]
		kmers = append(kmers, kmer)
	}
	return kmers
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba3a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read k
	scanner.Scan()
	k, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error parsing k:", err)
		return
	}

	// Read Text
	scanner.Scan()
	text := scanner.Text()

	// Generate k-mers
	kmers := generateKmers(text, k)

	// Sort k-mers lexicographically
	sort.Strings(kmers)

	// Remove duplicates
	uniqueKmers := make([]string, 0)
	for i, kmer := range kmers {
		if i == 0 || kmer != kmers[i-1] {
			uniqueKmers = append(uniqueKmers, kmer)
		}
	}

	// Print k-mers
	for _, kmer := range uniqueKmers {
		fmt.Println(kmer)
	}
}
