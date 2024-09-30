package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
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

func distanceBetweenPatternAndStrings(pattern string, dna []string) int {
	k := len(pattern)
	distance := 0

	for _, text := range dna {
		hammingDist := math.MaxInt32
		for i := 0; i <= len(text)-k; i++ {
			patternPrime := text[i : i+k]
			currentDistance := hammingDistance(pattern, patternPrime)
			if currentDistance < hammingDist {
				hammingDist = currentDistance
			}
		}
		distance += hammingDist
	}

	return distance
}

func main() {
	file, err := os.Open("rosalind_ba2h.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read Pattern
	scanner.Scan()
	pattern := scanner.Text()

	// Read DNA strings
	scanner.Scan()
	dna := strings.Split(scanner.Text(), " ")

	result := distanceBetweenPatternAndStrings(pattern, dna)
	fmt.Println(result)
}
