package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to calculate the GC content of a DNA string
func gcContent(dna string) float64 {
	gcCount := 0
	for _, nucleotide := range dna {
		if nucleotide == 'G' || nucleotide == 'C' {
			gcCount++
		}
	}
	return (float64(gcCount) / float64(len(dna))) * 100
}

// Function to parse FASTA format and calculate GC content
func highestGCContent(fileName string) (string, float64) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", 0.0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var maxID string
	var maxGC float64
	var currentID string
	var currentSeq strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		// If the line starts with '>', it's a new sequence
		if strings.HasPrefix(line, ">") {
			// If there is an existing sequence, calculate its GC content
			if currentID != "" {
				gc := gcContent(currentSeq.String())
				if gc > maxGC {
					maxGC = gc
					maxID = currentID
				}
			}
			// Extract the new sequence ID
			currentID = strings.TrimPrefix(line, ">")
			// Reset the sequence builder for the new sequence
			currentSeq.Reset()
		} else {
			// Accumulate the sequence data
			currentSeq.WriteString(line)
		}
	}

	// Compute GC content for the last sequence in the file
	if currentID != "" {
		gc := gcContent(currentSeq.String())
		if gc > maxGC {
			maxGC = gc
			maxID = currentID
		}
	}

	return maxID, maxGC
}

func main() {
	// File name for the input FASTA file
	fileName := "rosalind_gc.txt"

	// Get the ID and GC content of the sequence with the highest GC content
	id, gc := highestGCContent(fileName)

	// Print the result
	fmt.Printf("%s\n%.6f\n", id, gc)
}
