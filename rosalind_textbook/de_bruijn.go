package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func constructDeBruijnGraph(text string, k int) map[string][]string {
	graph := make(map[string][]string)

	for i := 0; i <= len(text)-k; i++ {
		kmer := text[i : i+k]
		prefix := kmer[:k-1]
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
	file, err := os.Open("rosalind_ba3d.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read k
	scanner.Scan()
	k, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Error parsing k:", err)
		return
	}

	// Read Text
	scanner.Scan()
	text := strings.TrimSpace(scanner.Text())

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Construct the de Bruijn graph
	graph := constructDeBruijnGraph(text, k)

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
