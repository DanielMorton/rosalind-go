package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// calculatePDist computes the p-distance between two DNA strings of equal length
func calculatePDist(s1, s2 string) float64 {
	if len(s1) != len(s2) {
		panic("Strings must be of equal length")
	}

	var diffCount float64
	for i := range s1 {
		if s1[i] != s2[i] {
			diffCount++
		}
	}

	return diffCount / float64(len(s1))
}

// readFastaFile reads a FASTA file and returns the DNA strings as a slice
func readFastaFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sequences []string
	var sb strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if sb.Len() > 0 {
				sequences = append(sequences, sb.String())
				sb.Reset()
			}
		} else {
			sb.WriteString(line)
		}
	}
	if sb.Len() > 0 {
		sequences = append(sequences, sb.String())
	}

	return sequences, scanner.Err()
}

func main() {
	// Read the sequences from the FASTA file
	sequences, err := readFastaFile("rosalind_pdst.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	n := len(sequences)
	if n == 0 {
		fmt.Println("No sequences found.")
		return
	}

	// Initialize the distance matrix
	D := make([][]float64, n)
	for i := range D {
		D[i] = make([]float64, n)
	}

	// Calculate the distance matrix
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			D[i][j] = calculatePDist(sequences[i], sequences[j])
		}
	}

	// Print the distance matrix with the required format
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%.5f", D[i][j])
		}
		fmt.Println()
	}
}
