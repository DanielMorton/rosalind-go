package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Function to count the occurrences of each 4-mer in the DNA string
func count4mers(dna string) map[string]int {
	// Initialize a map to store counts of each 4-mer
	counts := make(map[string]int)
	// Length of the DNA string
	n := len(dna)
	// Loop through all substrings of length 4
	for i := 0; i <= n-4; i++ {
		kmer := dna[i : i+4]
		counts[kmer]++
	}
	return counts
}

// Function to generate all possible 4-mers in lexicographic order
func generate4mers() []string {
	alphabet := "ACGT"
	var kmers []string
	// Generate all possible 4-mers
	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			for c := 0; c < 4; c++ {
				for d := 0; d < 4; d++ {
					kmer := string(alphabet[a]) + string(alphabet[b]) + string(alphabet[c]) + string(alphabet[d])
					kmers = append(kmers, kmer)
				}
			}
		}
	}
	return kmers
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_kmer.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var dna string

	// Read the DNA string from the file
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			continue // Skip header lines
		}
		dna += strings.TrimSpace(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Compute the 4-mer counts
	counts := count4mers(dna)
	// Generate all possible 4-mers in lexicographic order
	kmers := generate4mers()

	// Output the counts in lexicographic order
	for _, kmer := range kmers {
		fmt.Printf("%d ", counts[kmer])
	}
	fmt.Println()
}
