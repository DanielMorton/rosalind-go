package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func probabilityKmer(kmer string, profile [][]float64) float64 {
	probability := 1.0
	for i, nucleotide := range kmer {
		switch nucleotide {
		case 'A':
			probability *= profile[0][i]
		case 'C':
			probability *= profile[1][i]
		case 'G':
			probability *= profile[2][i]
		case 'T':
			probability *= profile[3][i]
		}
	}
	return probability
}

func profileMostProbableKmer(text string, k int, profile [][]float64) string {
	maxProbability := -1.0
	var mostProbableKmer string

	for i := 0; i <= len(text)-k; i++ {
		kmer := text[i : i+k]
		probability := probabilityKmer(kmer, profile)
		if probability > maxProbability {
			maxProbability = probability
			mostProbableKmer = kmer
		}
	}

	return mostProbableKmer
}

func main() {
	file, err := os.Open("rosalind_ba2c.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the text
	scanner.Scan()
	text := scanner.Text()

	// Read k
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	// Read the profile matrix
	profile := make([][]float64, 4)
	for i := 0; i < 4; i++ {
		scanner.Scan()
		line := scanner.Text()
		probabilities := strings.Split(line, " ")
		profile[i] = make([]float64, k)
		for j, prob := range probabilities {
			profile[i][j], _ = strconv.ParseFloat(prob, 64)
		}
	}

	result := profileMostProbableKmer(text, k, profile)
	fmt.Println(result)
}
