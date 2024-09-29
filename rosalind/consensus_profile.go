package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("rosalind_cons.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the sequences from the file
	scanner := bufio.NewScanner(file)
	var sequences []string
	var currentSequence strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			// If we hit a new header and have a current sequence, add it to sequences
			if currentSequence.Len() > 0 {
				sequences = append(sequences, currentSequence.String())
				currentSequence.Reset()
			}
		} else {
			// Append the line to the current sequence
			currentSequence.WriteString(line)
		}
	}

	// Append the last sequence
	if currentSequence.Len() > 0 {
		sequences = append(sequences, currentSequence.String())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(sequences) == 0 {
		fmt.Println("No sequences found")
		return
	}

	// Determine the length of the sequences
	n := len(sequences[0])

	// Initialize the profile matrix
	profile := [4][]int{
		make([]int, n), // A
		make([]int, n), // C
		make([]int, n), // G
		make([]int, n), // T
	}

	// Fill the profile matrix
	for _, seq := range sequences {
		for j := 0; j < n; j++ {
			switch seq[j] {
			case 'A':
				profile[0][j]++
			case 'C':
				profile[1][j]++
			case 'G':
				profile[2][j]++
			case 'T':
				profile[3][j]++
			}
		}
	}

	// Build the consensus string
	var consensus strings.Builder
	for j := 0; j < n; j++ {
		maxCount := -1
		maxChar := ' '
		for i, char := range []rune{'A', 'C', 'G', 'T'} {
			if profile[i][j] > maxCount {
				maxCount = profile[i][j]
				maxChar = char
			}
		}
		consensus.WriteRune(maxChar)
	}

	// Print the consensus string
	fmt.Println(consensus.String())

	// Print the profile matrix
	for i, char := range []rune{'A', 'C', 'G', 'T'} {
		fmt.Printf("%c: ", char)
		for _, count := range profile[i] {
			fmt.Printf("%d ", count)
		}
		fmt.Println()
	}
}
