package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph map[string][]string

func constructDeBruijnGraph(kmers []string) Graph {
	graph := make(Graph)
	for _, kmer := range kmers {
		prefix := kmer[:len(kmer)-1]
		suffix := kmer[1:]
		graph[prefix] = append(graph[prefix], suffix)
	}
	return graph
}

func findStartNode(graph Graph) string {
	inDegree := make(map[string]int)
	for _, neighbors := range graph {
		for _, neighbor := range neighbors {
			inDegree[neighbor]++
		}
	}

	for node := range graph {
		if len(graph[node]) > inDegree[node] {
			return node
		}
	}

	// If no start node found, return any node
	for node := range graph {
		return node
	}
	return ""
}

func findEulerianPath(graph Graph) []string {
	start := findStartNode(graph)
	stack := []string{start}
	path := []string{}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		if len(graph[current]) > 0 {
			next := graph[current][0]
			stack = append(stack, next)
			graph[current] = graph[current][1:]
		} else {
			path = append([]string{current}, path...)
			stack = stack[:len(stack)-1]
		}
	}

	return path
}

func reconstructString(path []string) string {
	if len(path) == 0 {
		return ""
	}
	result := path[0]
	for i := 1; i < len(path); i++ {
		result += string(path[i][len(path[i])-1])
	}
	return result
}

func main() {
	file, err := os.Open("rosalind_ba3h.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read k
	scanner.Scan()
	k, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	// Read k-mers
	var kmers []string
	for scanner.Scan() {
		kmer := strings.TrimSpace(scanner.Text())
		if kmer != "" {
			if len(kmer) != k {
				fmt.Printf("Warning: kmer '%s' length does not match k=%d\n", kmer, k)
			}
			kmers = append(kmers, kmer)
		}
	}

	// Construct de Bruijn graph
	graph := constructDeBruijnGraph(kmers)

	// Find Eulerian path
	path := findEulerianPath(graph)

	// Reconstruct string
	result := reconstructString(path)

	// Verify result
	if len(result) != len(kmers[0])+len(kmers)-1 {
		fmt.Printf("Warning: Reconstructed string length (%d) doesn't match expected length (%d)\n",
			len(result), len(kmers[0])+len(kmers)-1)
	}

	// Print result
	fmt.Println(result)
}
