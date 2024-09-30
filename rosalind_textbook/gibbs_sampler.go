package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func createProfile(motifs []string, excludeIndex int) [][]float64 {
	k := len(motifs[0])
	profile := make([][]float64, 4)
	for i := range profile {
		profile[i] = make([]float64, k)
	}

	for i := 0; i < k; i++ {
		counts := map[byte]float64{'A': 1, 'C': 1, 'G': 1, 'T': 1} // Pseudocounts
		for j, motif := range motifs {
			if j != excludeIndex {
				counts[motif[i]]++
			}
		}
		total := float64(len(motifs) - 1 + 4) // -1 for excluded motif, +4 for pseudocounts
		profile[0][i] = counts['A'] / total
		profile[1][i] = counts['C'] / total
		profile[2][i] = counts['G'] / total
		profile[3][i] = counts['T'] / total
	}

	return profile
}

func profileRandomlyGeneratedKmer(text string, k int, profile [][]float64) string {
	probabilities := make([]float64, len(text)-k+1)
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
		probabilities[i] = probability
	}

	// Choose a k-mer based on the probabilities
	total := 0.0
	for _, p := range probabilities {
		total += p
	}
	r := rand.Float64() * total
	sum := 0.0
	for i, p := range probabilities {
		sum += p
		if sum > r {
			return text[i : i+k]
		}
	}
	return text[:k] // This should never happen, but just in case
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

func randomMotifs(dna []string, k int) []string {
	motifs := make([]string, len(dna))
	for i, seq := range dna {
		start := rand.Intn(len(seq) - k + 1)
		motifs[i] = seq[start : start+k]
	}
	return motifs
}

func gibbsSampler(dna []string, k, t, N int) []string {
	motifs := randomMotifs(dna, k)
	bestMotifs := make([]string, len(motifs))
	copy(bestMotifs, motifs)

	for j := 0; j < N; j++ {
		i := rand.Intn(t)
		profile := createProfile(motifs, i)
		motifs[i] = profileRandomlyGeneratedKmer(dna[i], k, profile)
		if score(motifs) < score(bestMotifs) {
			copy(bestMotifs, motifs)
		}
	}

	return bestMotifs
}

func main() {
	file, err := os.Open("rosalind_ba2g.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read k, t, and N
	scanner.Scan()
	params := strings.Split(scanner.Text(), " ")
	k, _ := strconv.Atoi(params[0])
	t, _ := strconv.Atoi(params[1])
	N, _ := strconv.Atoi(params[2])

	// Read DNA strings
	dna := make([]string, t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		dna[i] = scanner.Text()
	}

	rand.Seed(time.Now().UnixNano())

	bestScore := int(^uint(0) >> 1) // Max int
	var bestMotifs []string

	for i := 0; i < 20; i++ {
		motifs := gibbsSampler(dna, k, t, N)
		currentScore := score(motifs)
		if currentScore < bestScore {
			bestScore = currentScore
			bestMotifs = motifs
		}
	}

	for _, motif := range bestMotifs {
		fmt.Println(motif)
	}
}
