package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func constructDeBruijnGraph(patterns []string) map[string][]string {
	graph := make(map[string][]string)

	for _, kmer := range patterns {
		prefix := kmer[:len(kmer)-1]
		suffix := kmer[1:]

		if _, exists := graph[prefix]; !exists {
			graph[prefix] = []string{}
		}
		graph[prefix] = append(graph[prefix], suffix)
	}

	// Sort the adjacency list for each node
	for node := range graph {
		sort.Strings(graph[node])
	}

	return graph
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba3e.txt")
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

	// Construct the de Bruijn graph
	graph := constructDeBruijnGraph(patterns)

	// Print the adjacency list
	var nodes []string
	for node := range graph {
		nodes = append(nodes, node)
	}
	sort.Strings(nodes)

	for _, node := range nodes {
		neighbors := graph[node]
		fmt.Printf("%s -> %s\n", node, strings.Join(neighbors, ","))
	}
}
