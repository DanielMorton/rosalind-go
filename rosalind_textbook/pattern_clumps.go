package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// findClumps returns all distinct k-mers forming (L, t)-clumps in Genome
func findClumps(genome string, k, L, t int) []string {
	clumps := make(map[string]bool)
	for i := 0; i <= len(genome)-L; i++ {
		window := genome[i : i+L]
		kmerCounts := make(map[string]int)

		// Count k-mers in the current window
		for j := 0; j <= L-k; j++ {
			kmer := window[j : j+k]
			kmerCounts[kmer]++
			if kmerCounts[kmer] >= t {
				clumps[kmer] = true
			}
		}
	}

	// Convert map keys to slice
	var result []string
	for kmer := range clumps {
		result = append(result, kmer)
	}
	return result
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba1e.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read input from file
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	genome := strings.TrimSpace(scanner.Text())
	scanner.Scan()
	params := strings.Fields(scanner.Text())

	if len(params) != 3 {
		fmt.Println("Invalid input format")
		return
	}

	k, _ := strconv.Atoi(params[0])
	L, _ := strconv.Atoi(params[1])
	t, _ := strconv.Atoi(params[2])

	// Find clumps
	clumps := findClumps(genome, k, L, t)

	// Print the result
	fmt.Println(strings.Join(clumps, " "))
}
