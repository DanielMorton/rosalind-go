package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// numberToPattern converts a number to its corresponding DNA string of length k
func numberToPattern(index, k int) string {
	nucleotides := []byte{'A', 'C', 'G', 'T'}
	pattern := make([]byte, k)

	for i := k - 1; i >= 0; i-- {
		remainder := index % 4
		pattern[i] = nucleotides[remainder]
		index = index / 4
	}

	return string(pattern)
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba1m.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read input from file
	scanner := bufio.NewScanner(file)

	// Read index
	scanner.Scan()
	index, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Error parsing index:", err)
		return
	}

	// Read k
	scanner.Scan()
	k, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Error parsing k:", err)
		return
	}

	// Convert number to pattern
	pattern := numberToPattern(index, k)

	// Print the result
	fmt.Println(pattern)
}
