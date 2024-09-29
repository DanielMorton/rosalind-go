package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReverseComplement returns the reverse complement of a DNA string.
func ReverseComplement(dna string) string {
	var complement = map[byte]byte{'A': 'T', 'T': 'A', 'C': 'G', 'G': 'C'}
	var revComp strings.Builder
	for i := len(dna) - 1; i >= 0; i-- {
		revComp.WriteByte(complement[dna[i]])
	}
	return revComp.String()
}

// HammingDistance calculates the Hamming distance between two strings.
func HammingDistance(a, b string) int {
	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance
}

func main() {
	// Read the input file
	file, err := os.Open("rosalind_corr.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var reads []string
	readCount := make(map[string]int)

	// Parse the FASTA format file
	var currentRead string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if currentRead != "" {
				reads = append(reads, currentRead)
				readCount[currentRead]++
				revComp := ReverseComplement(currentRead)
				readCount[revComp]++
			}
			currentRead = ""
		} else {
			currentRead += line
		}
	}
	// Append the last read
	if currentRead != "" {
		reads = append(reads, currentRead)
		readCount[currentRead]++
		revComp := ReverseComplement(currentRead)
		readCount[revComp]++
	}

	// Identify correct and incorrect reads
	var correctReads, incorrectReads []string
	for _, read := range reads {
		if readCount[read] > 1 {
			correctReads = append(correctReads, read)
		} else {
			incorrectReads = append(incorrectReads, read)
		}
	}

	// Find corrections
	var corrections []string
	for _, incorrectRead := range incorrectReads {
		for _, correctRead := range correctReads {
			if HammingDistance(incorrectRead, correctRead) == 1 {
				corrections = append(corrections, fmt.Sprintf("%s->%s", incorrectRead, correctRead))
				break
			}
			if HammingDistance(incorrectRead, ReverseComplement(correctRead)) == 1 {
				corrections = append(corrections, fmt.Sprintf("%s->%s", incorrectRead, correctRead))
				break
			}
		}
	}

	// Output the corrections
	for _, correction := range corrections {
		fmt.Println(correction)
	}
}
