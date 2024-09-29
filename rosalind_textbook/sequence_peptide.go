package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AminoAcid struct {
	Mass   int
	Symbol string
}

var aminoAcids = []AminoAcid{
	{71, "A"}, {103, "C"}, {115, "D"}, {129, "E"}, {147, "F"},
	{57, "G"}, {137, "H"}, {113, "I"}, {128, "K"}, {113, "L"},
	{131, "M"}, {114, "N"}, {97, "P"}, {128, "Q"}, {156, "R"},
	{87, "S"}, {101, "T"}, {99, "V"}, {186, "W"}, {163, "Y"},
	//{4, "X"}, {5, "Z"}, // Imaginary amino acids X and Z
}

type Node struct {
	Weight   int
	MaxScore int
	BestPrev int
	BestAA   string
}

func main() {
	spectralVector := readSpectralVector("rosalind_ba11e.txt")
	result := findOptimalPeptide(spectralVector)
	fmt.Println(result)
}

func readSpectralVector(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := strings.TrimSpace(scanner.Text())

	var vector []int
	vector = append(vector, 0) // Add a 0 at index 0 to align with 1-based indexing
	for _, s := range strings.Fields(line) {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Invalid input: %s", s))
		}
		vector = append(vector, num)
	}
	return vector
}

func findOptimalPeptide(spectralVector []int) string {
	m := len(spectralVector)
	nodes := make([]Node, m)

	// Initialize nodes
	nodes[0] = Node{Weight: 0, MaxScore: 0, BestPrev: -1, BestAA: ""}
	for i := 1; i < m; i++ {
		nodes[i] = Node{Weight: spectralVector[i], MaxScore: -1000000, BestPrev: -1, BestAA: ""}
	}

	// Dynamic programming to find the path with maximum total weight
	for i := 1; i < m; i++ {
		for _, aa := range aminoAcids {
			if i-aa.Mass >= 0 {
				score := nodes[i-aa.Mass].MaxScore + nodes[i].Weight
				if score > nodes[i].MaxScore {
					nodes[i].MaxScore = score
					nodes[i].BestPrev = i - aa.Mass
					nodes[i].BestAA = aa.Symbol
				} else if score == nodes[i].MaxScore && aa.Symbol == "Z" {
					// Prefer 'Z' over other amino acids when scores are equal
					nodes[i].BestPrev = i - aa.Mass
					nodes[i].BestAA = aa.Symbol
				}
			}
		}
	}

	// Reconstruct the peptide
	var peptide strings.Builder
	for i := m - 1; i > 0; {
		peptide.WriteString(nodes[i].BestAA)
		i = nodes[i].BestPrev
	}

	return reverseString(peptide.String())
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
