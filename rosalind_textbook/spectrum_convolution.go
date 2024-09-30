package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func computeConvolution(spectrum []int) []int {
	var convolution []int
	for i := 0; i < len(spectrum); i++ {
		for j := 0; j < len(spectrum); j++ {
			diff := spectrum[j] - spectrum[i]
			if diff > 0 {
				convolution = append(convolution, diff)
			}
		}
	}
	return convolution
}

func sortConvolution(convolution []int) []int {
	// Count frequencies
	freqMap := make(map[int]int)
	for _, mass := range convolution {
		freqMap[mass]++
	}

	// Create a slice of unique masses
	var uniqueMasses []int
	for mass := range freqMap {
		uniqueMasses = append(uniqueMasses, mass)
	}

	// Sort unique masses by frequency (descending) and then by mass (ascending)
	sort.Slice(uniqueMasses, func(i, j int) bool {
		if freqMap[uniqueMasses[i]] == freqMap[uniqueMasses[j]] {
			return uniqueMasses[i] < uniqueMasses[j]
		}
		return freqMap[uniqueMasses[i]] > freqMap[uniqueMasses[j]]
	})

	// Create the final sorted convolution
	var sortedConvolution []int
	for _, mass := range uniqueMasses {
		for i := 0; i < freqMap[mass]; i++ {
			sortedConvolution = append(sortedConvolution, mass)
		}
	}

	return sortedConvolution
}

func main() {
	file, err := os.Open("rosalind_ba4h.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	spectrumStr := strings.Fields(scanner.Text())

	spectrum := make([]int, len(spectrumStr))
	for i, s := range spectrumStr {
		spectrum[i], _ = strconv.Atoi(s)
	}

	convolution := computeConvolution(spectrum)
	sortedConvolution := sortConvolution(convolution)

	// Print the result
	for i, mass := range sortedConvolution {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(mass)
	}
	fmt.Println()

	// Print the length of the output
	//fmt.Printf("\nOutput length: %d\n", len(sortedConvolution))
}
