package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input from file
	file, err := os.Open("rosalind_cstr.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var dnaStrings []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dnaStrings = append(dnaStrings, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Process the DNA strings
	result := processStrings(dnaStrings)

	// Print the result
	for _, row := range result {
		fmt.Println(row)
	}
}

func processStrings(dnaStrings []string) []string {
	if len(dnaStrings) == 0 {
		return nil
	}

	length := len(dnaStrings[0])
	var result []string

	for i := 0; i < length; i++ {
		counts := make(map[byte]int)
		for _, dna := range dnaStrings {
			counts[dna[i]]++
		}

		if len(counts) == 2 {
			var char string
			var nucleotides []byte
			for nuc := range counts {
				nucleotides = append(nucleotides, nuc)
			}

			// Check if both nucleotides appear at least twice
			if counts[nucleotides[0]] >= 2 && counts[nucleotides[1]] >= 2 {
				for _, dna := range dnaStrings {
					if dna[i] == nucleotides[0] {
						char += "1"
					} else {
						char += "0"
					}
				}
				result = append(result, char)
			}
		}
	}

	return result
}
