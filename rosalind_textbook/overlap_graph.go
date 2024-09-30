package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func suffix(s string) string {
	return s[1:]
}

func prefix(s string) string {
	return s[:len(s)-1]
}

func constructOverlapGraph(patterns []string) map[string][]string {
	graph := make(map[string][]string)

	for _, pattern := range patterns {
		graph[pattern] = []string{}
	}

	for _, pattern1 := range patterns {
		for _, pattern2 := range patterns {
			if pattern1 != pattern2 && suffix(pattern1) == prefix(pattern2) {
				graph[pattern1] = append(graph[pattern1], pattern2)
			}
		}
	}

	return graph
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba3c.txt")
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

	// Construct the overlap graph
	graph := constructOverlapGraph(patterns)

	// Print the adjacency list
	var keys []string
	for k := range graph {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, pattern := range keys {
		neighbors := graph[pattern]
		if len(neighbors) > 0 {
			sort.Strings(neighbors)
			for _, neighbor := range neighbors {
				fmt.Printf("%s -> %s\n", pattern, neighbor)
			}
		}
	}
}
