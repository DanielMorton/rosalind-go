package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadDNA reads the DNA sequence from the input file, handling multiple lines.
func ReadDNA(filename string) (string, error) {
	var sb strings.Builder
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, ">") { // Skip FASTA header lines
			sb.WriteString(line)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return sb.String(), nil
}

// ReverseComplement returns the reverse complement of a DNA string.
func ReverseComplement(dna string) string {
	complement := map[rune]rune{
		'A': 'T',
		'T': 'A',
		'C': 'G',
		'G': 'C',
	}
	var revComp strings.Builder
	for i := len(dna) - 1; i >= 0; i-- {
		revComp.WriteRune(complement[rune(dna[i])])
	}
	return revComp.String()
}

// FindReversePalindromes finds all reverse palindromes of lengths between minLen and maxLen.
func FindReversePalindromes(dna string, minLen, maxLen int) []struct {
	position int
	length   int
} {
	var results []struct {
		position int
		length   int
	}

	length := len(dna)
	for start := 0; start < length; start++ {
		for end := start + minLen; end <= start+maxLen && end <= length; end++ {
			substr := dna[start:end]
			if substr == ReverseComplement(substr) {
				results = append(results, struct {
					position int
					length   int
				}{
					position: start + 1, // 1-based index
					length:   end - start,
				})
			}
		}
	}
	return results
}

func main() {
	dna, err := ReadDNA("rosalind_revp.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	results := FindReversePalindromes(dna, 4, 12)
	for _, result := range results {
		fmt.Printf("%d %d\n", result.position, result.length)
	}
}
