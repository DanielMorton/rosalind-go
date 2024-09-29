package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Function to read FASTA format data
func readFASTA(filename string) map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sequences := make(map[string]string)
	var currentLabel string
	var currentSeq strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			// Save the current sequence if there's one being built
			if currentLabel != "" {
				sequences[currentLabel] = currentSeq.String()
			}
			// Start a new label and reset sequence builder
			currentLabel = line[1:] // Remove '>'
			currentSeq.Reset()
		} else {
			// Build the sequence
			currentSeq.WriteString(line)
		}
	}
	// Add the last sequence
	if currentLabel != "" {
		sequences[currentLabel] = currentSeq.String()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sequences
}

// Function to construct the overlap graph adjacency list
func constructOverlapGraph(sequences map[string]string, k int) {
	for label1, seq1 := range sequences {
		suffix := seq1[len(seq1)-k:] // suffix of length k
		for label2, seq2 := range sequences {
			if label1 != label2 {
				prefix := seq2[:k] // prefix of length k
				if suffix == prefix {
					fmt.Println(label1, label2)
				}
			}
		}
	}
}

func main() {
	// Read the FASTA sequences from file
	sequences := readFASTA("rosalind_grph.txt")

	// Construct and print the adjacency list for overlap graph O3
	constructOverlapGraph(sequences, 3)
}
