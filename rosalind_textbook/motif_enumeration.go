package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func neighbors(pattern string, d int) []string {
	if d == 0 {
		return []string{pattern}
	}
	if len(pattern) == 1 {
		return []string{"A", "C", "G", "T"}
	}
	neighborhood := []string{}
	suffixNeighbors := neighbors(pattern[1:], d)
	for _, text := range suffixNeighbors {
		if hammingDistance(pattern[1:], text) < d {
			for _, nucleotide := range []string{"A", "C", "G", "T"} {
				neighborhood = append(neighborhood, nucleotide+text)
			}
		} else {
			neighborhood = append(neighborhood, string(pattern[0])+text)
		}
	}
	return neighborhood
}

func motifEnumeration(dna []string, k, d int) []string {
	patterns := make(map[string]bool)
	for i, text := range dna {
		for j := 0; j <= len(text)-k; j++ {
			pattern := text[j : j+k]
			neighborhood := neighbors(pattern, d)
			for _, neighbor := range neighborhood {
				isMotif := true
				for m, seq := range dna {
					if i == m {
						continue
					}
					found := false
					for n := 0; n <= len(seq)-k; n++ {
						if hammingDistance(neighbor, seq[n:n+k]) <= d {
							found = true
							break
						}
					}
					if !found {
						isMotif = false
						break
					}
				}
				if isMotif {
					patterns[neighbor] = true
				}
			}
		}
	}
	result := []string{}
	for pattern := range patterns {
		result = append(result, pattern)
	}
	return result
}

func main() {
	file, err := os.Open("rosalind_ba2a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	params := strings.Split(scanner.Text(), " ")
	k, _ := strconv.Atoi(params[0])
	d, _ := strconv.Atoi(params[1])

	dna := []string{}
	for scanner.Scan() {
		dna = append(dna, scanner.Text())
	}

	result := motifEnumeration(dna, k, d)
	fmt.Println(strings.Join(result, " "))
}
