package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// patternToNumber converts a DNA string to a numerical representation
func patternToNumber(pattern string) int {
	nucleotideMap := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	result := 0
	for i := 0; i < len(pattern); i++ {
		result = 4*result + nucleotideMap[pattern[i]]
	}
	return result
}

// computeFrequencyArray generates the frequency array of k-mers in text
func computeFrequencyArray(text string, k int) []int {
	frequencyArray := make([]int, int(math.Pow(4, float64(k))))

	for i := 0; i <= len(text)-k; i++ {
		pattern := text[i : i+k]
		j := patternToNumber(pattern)
		frequencyArray[j]++
	}

	return frequencyArray
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba1k.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read input from file
	scanner := bufio.NewScanner(file)

	// Read Text
	scanner.Scan()
	text := strings.TrimSpace(scanner.Text())

	// Read k
	scanner.Scan()
	k, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Error parsing k:", err)
		return
	}

	// Compute frequency array
	frequencyArray := computeFrequencyArray(text, k)

	// Print the result
	for i, freq := range frequencyArray {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(freq)
	}
	fmt.Println()
}
