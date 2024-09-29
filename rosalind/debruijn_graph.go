package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func reverseComplement(s string) string {
	complement := map[rune]rune{'A': 'T', 'T': 'A', 'C': 'G', 'G': 'C'}
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = complement[runes[j]], complement[runes[i]]
	}
	return string(runes)
}

func generateKMers(s string, k int) []string {
	result := make([]string, 0)
	for i := 0; i <= len(s)-k; i++ {
		result = append(result, s[i:i+k])
	}
	return result
}

func main() {
	file, err := os.Open("rosalind_dbru.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	set := make(map[string]bool)
	k := 0

	for scanner.Scan() {
		seq := strings.TrimSpace(scanner.Text())
		if seq != "" {
			set[seq] = true
			set[reverseComplement(seq)] = true
			if k == 0 {
				k = len(seq) - 1 // Set k only once, using the first non-empty sequence
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if k == 0 {
		fmt.Println("No valid sequences found in the file")
		return
	}

	edges := make(map[string]bool)

	for seq := range set {
		kmers := generateKMers(seq, k)
		for i := 0; i < len(kmers)-1; i++ {
			edge := fmt.Sprintf("(%s, %s)", kmers[i], kmers[i+1])
			edges[edge] = true
		}
	}

	sortedEdges := make([]string, 0, len(edges))
	for edge := range edges {
		sortedEdges = append(sortedEdges, edge)
	}
	sort.Strings(sortedEdges)

	for _, edge := range sortedEdges {
		fmt.Println(edge)
	}
}
