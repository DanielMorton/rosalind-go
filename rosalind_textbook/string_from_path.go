package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func stringSpelledByGenomePath(patterns []string) string {
	if len(patterns) == 0 {
		return ""
	}

	result := patterns[0]
	k := len(patterns[0])

	for i := 1; i < len(patterns); i++ {
		result += string(patterns[i][k-1])
	}

	return result
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba3b.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read patterns
	var patterns []string
	for scanner.Scan() {
		pattern := strings.TrimSpace(scanner.Text())
		if pattern != "" {
			patterns = append(patterns, pattern)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Generate the string spelled by the genome path
	result := stringSpelledByGenomePath(patterns)

	// Print the result
	fmt.Println(result)
}
