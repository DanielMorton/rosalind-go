package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var aminoAcidMasses = []int{
	57, 71, 87, 97, 99, 101, 103, 113, 114, 115,
	128, 129, 131, 137, 147, 156, 163, 186,
	4, 5, // Imaginary amino acids X and Z
}

func main() {
	spectrum, threshold, maxScore := readInput("rosalind_ba11h.txt")
	result := sizeOfSpectralDictionary(spectrum, threshold, maxScore)
	fmt.Println(result.String())
}

func readInput(filename string) ([]int, int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read spectral vector
	scanner.Scan()
	spectrumStr := strings.Fields(scanner.Text())
	spectrum := make([]int, len(spectrumStr))
	for i, s := range spectrumStr {
		spectrum[i], _ = strconv.Atoi(s)
	}

	// Read threshold
	scanner.Scan()
	threshold, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	// Read max_score
	scanner.Scan()
	maxScore, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	return spectrum, threshold, maxScore
}

func sizeOfSpectralDictionary(spectrum []int, threshold, maxScore int) *big.Int {
	m := len(spectrum)
	minScore := -1000 // Assume a large negative value as the minimum possible score
	scoreRange := maxScore - minScore + 1

	// Initialize the DP table with big.Int
	size := make([][]*big.Int, m+1)
	for i := range size {
		size[i] = make([]*big.Int, scoreRange)
		for j := range size[i] {
			size[i][j] = big.NewInt(0)
		}
	}

	// Base case: empty peptide
	size[0][-minScore].SetInt64(1)

	// Fill the DP table
	for i := 1; i <= m; i++ {
		for _, mass := range aminoAcidMasses {
			if i >= mass {
				for s := minScore; s <= maxScore; s++ {
					prevScore := s - spectrum[i-1]
					if prevScore >= minScore && prevScore <= maxScore {
						size[i][s-minScore].Add(size[i][s-minScore], size[i-mass][prevScore-minScore])
					}
				}
			}
		}
	}

	// Sum up all sizes for scores >= threshold and <= maxScore
	totalSize := big.NewInt(0)
	for s := threshold; s <= maxScore; s++ {
		totalSize.Add(totalSize, size[m][s-minScore])
	}

	return totalSize
}
