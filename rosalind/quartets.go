package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func parsePartialCharacterTable(lines []string) ([]string, [][]rune) {
	taxa := strings.Fields(lines[0])
	table := make([][]rune, len(lines)-1)
	for i, line := range lines[1:] {
		table[i] = []rune(line)
	}
	return taxa, table
}

func generateQuartets(taxa []string, row []rune) []string {
	var set0, set1 []string
	for i, char := range row {
		if char == '0' {
			set0 = append(set0, taxa[i])
		} else if char == '1' {
			set1 = append(set1, taxa[i])
		}
	}

	var quartets []string
	for i := 0; i < len(set0); i++ {
		for j := i + 1; j < len(set0); j++ {
			for k := 0; k < len(set1); k++ {
				for l := k + 1; l < len(set1); l++ {
					pair1 := []string{set0[i], set0[j]}
					pair2 := []string{set1[k], set1[l]}
					sort.Strings(pair1)
					sort.Strings(pair2)
					if pair1[0] > pair2[0] {
						pair1, pair2 = pair2, pair1
					}
					quartet := fmt.Sprintf("{%s, %s} {%s, %s}", pair1[0], pair1[1], pair2[0], pair2[1])
					quartets = append(quartets, quartet)
				}
			}
		}
	}
	return quartets
}

func main() {
	file, err := os.Open("rosalind_qrt.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	taxa, table := parsePartialCharacterTable(lines)

	uniqueQuartets := make(map[string]bool)
	for _, row := range table {
		quartets := generateQuartets(taxa, row)
		for _, quartet := range quartets {
			uniqueQuartets[quartet] = true
		}
	}

	// Convert map to slice for sorting
	quartetSlice := make([]string, 0, len(uniqueQuartets))
	for quartet := range uniqueQuartets {
		quartetSlice = append(quartetSlice, quartet)
	}

	// Sort the quartets
	sort.Strings(quartetSlice)

	// Print sorted quartets
	for _, quartet := range quartetSlice {
		fmt.Println(quartet)
	}
}
