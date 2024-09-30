package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reconstructString(pairs []string, k, d int) string {
	if len(pairs) == 0 {
		return ""
	}

	// Extract the first and second parts of each pair
	firstParts := make([]string, len(pairs))
	secondParts := make([]string, len(pairs))
	for i, pair := range pairs {
		parts := strings.Split(pair, "|")
		firstParts[i] = parts[0]
		secondParts[i] = parts[1]
	}

	// Reconstruct the prefix string
	prefix := firstParts[0]
	for i := 1; i < len(firstParts); i++ {
		prefix += string(firstParts[i][k-1])
	}

	// Reconstruct the suffix string
	suffix := secondParts[0]
	for i := 1; i < len(secondParts); i++ {
		suffix += string(secondParts[i][k-1])
	}

	// Combine prefix and suffix
	return prefix + suffix[len(suffix)-d-k:]
}

func main() {
	file, err := os.Open("rosalind_ba3l.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read k and d
	scanner.Scan()
	var k, d int
	fmt.Sscanf(scanner.Text(), "%d %d", &k, &d)

	// Read (k,d)-mers
	var pairs []string
	for scanner.Scan() {
		pairs = append(pairs, scanner.Text())
	}

	// Reconstruct the string
	result := reconstructString(pairs, k, d)

	// Print result
	fmt.Println(result)
}
