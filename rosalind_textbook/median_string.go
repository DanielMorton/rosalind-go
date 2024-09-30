package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func hammingDistance(s1, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}
	distance := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			distance++
		}
	}
	return distance
}

func distancePatternString(pattern string, text string) int {
	k := len(pattern)
	minDistance := math.MaxInt32
	for i := 0; i <= len(text)-k; i++ {
		kmer := text[i : i+k]
		distance := hammingDistance(pattern, kmer)
		if distance < minDistance {
			minDistance = distance
		}
	}
	return minDistance
}

func distancePatternStrings(pattern string, dna []string) int {
	distance := 0
	for _, text := range dna {
		distance += distancePatternString(pattern, text)
	}
	return distance
}

func generateAllKmers(k int) []string {
	if k == 0 {
		return []string{""}
	}
	prevKmers := generateAllKmers(k - 1)
	kmers := []string{}
	for _, prevKmer := range prevKmers {
		for _, nucleotide := range []string{"A", "C", "G", "T"} {
			kmers = append(kmers, prevKmer+nucleotide)
		}
	}
	return kmers
}

func medianString(dna []string, k int) string {
	distance := math.MaxInt32
	var medianPattern string
	allKmers := generateAllKmers(k)

	for _, pattern := range allKmers {
		currentDistance := distancePatternStrings(pattern, dna)
		if currentDistance < distance {
			distance = currentDistance
			medianPattern = pattern
		}
	}
	return medianPattern
}

func main() {
	file, err := os.Open("rosalind_ba2b.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	dna := []string{}
	for scanner.Scan() {
		dna = append(dna, scanner.Text())
	}

	result := medianString(dna, k)
	fmt.Println(result)
}
