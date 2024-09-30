package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func profileMostProbableKmer(text string, k int, profile [][]float64) string {
	maxProbability := -1.0
	var mostProbableKmer string

	for i := 0; i <= len(text)-k; i++ {
		kmer := text[i : i+k]
		probability := 1.0
		for j, nucleotide := range kmer {
			switch nucleotide {
			case 'A':
				probability *= profile[0][j]
			case 'C':
				probability *= profile[1][j]
			case 'G':
				probability *= profile[2][j]
			case 'T':
				probability *= profile[3][j]
			}
		}
		if probability > maxProbability {
			maxProbability = probability
			mostProbableKmer = kmer
		}
	}

	return mostProbableKmer
}

func createProfile(motifs []string) [][]float64 {
	k := len(motifs[0])
	profile := make([][]float64, 4)
	for i := range profile {
		profile[i] = make([]float64, k)
	}

	for i := 0; i < k; i++ {
		counts := map[byte]float64{'A': 1, 'C': 1, 'G': 1, 'T': 1} // Pseudocounts
		for _, motif := range motifs {
			counts[motif[i]]++
		}
		total := float64(len(motifs) + 4) // +4 for pseudocounts
		profile[0][i] = counts['A'] / total
		profile[1][i] = counts['C'] / total
		profile[2][i] = counts['G'] / total
		profile[3][i] = counts['T'] / total
	}

	return profile
}

func score(motifs []string) int {
	score := 0
	for i := 0; i < len(motifs[0]); i++ {
		counts := map[byte]int{'A': 0, 'C': 0, 'G': 0, 'T': 0}
		for _, motif := range motifs {
			counts[motif[i]]++
		}
		maxCount := 0
		for _, count := range counts {
			if count > maxCount {
				maxCount = count
			}
		}
		score += len(motifs) - maxCount
	}
	return score
}

func greedyMotifSearch(dna []string, k, t int) []string {
	bestMotifs := make([]string, t)
	for i, seq := range dna {
		bestMotifs[i] = seq[:k]
	}

	for i := 0; i <= len(dna[0])-k; i++ {
		motifs := make([]string, t)
		motifs[0] = dna[0][i : i+k]

		for j := 1; j < t; j++ {
			profile := createProfile(motifs[:j])
			motifs[j] = profileMostProbableKmer(dna[j], k, profile)
		}

		if score(motifs) < score(bestMotifs) {
			bestMotifs = motifs
		}
	}

	return bestMotifs
}

func main() {
	file, err := os.Open("rosalind_ba2e.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read k and t
	scanner.Scan()
	params := strings.Split(scanner.Text(), " ")
	k, _ := strconv.Atoi(params[0])
	t, _ := strconv.Atoi(params[1])

	// Read DNA strings
	dna := make([]string, t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		dna[i] = scanner.Text()
	}

	bestMotifs := greedyMotifSearch(dna, k, t)
	for _, motif := range bestMotifs {
		fmt.Println(motif)
	}
}
