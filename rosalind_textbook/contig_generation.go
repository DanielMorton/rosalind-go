package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Graph map[string][]string

func constructDeBruijnGraph(patterns []string) Graph {
	graph := make(Graph)
	for _, pattern := range patterns {
		prefix := pattern[:len(pattern)-1]
		suffix := pattern[1:]
		graph[prefix] = append(graph[prefix], suffix)
	}
	return graph
}

func inDegree(graph Graph, node string) int {
	count := 0
	for _, neighbors := range graph {
		for _, neighbor := range neighbors {
			if neighbor == node {
				count++
			}
		}
	}
	return count
}

func findMaximalNonBranchingPaths(graph Graph) []string {
	paths := []string{}
	visited := make(map[string]bool)

	for v := range graph {
		if !(inDegree(graph, v) == 1 && len(graph[v]) == 1) {
			if len(graph[v]) > 0 {
				for _, w := range graph[v] {
					nonBranchingPath := v + w[len(w)-1:]
					visited[v] = true
					u := w
					for inDegree(graph, u) == 1 && len(graph[u]) == 1 {
						visited[u] = true
						nonBranchingPath += graph[u][0][len(graph[u][0])-1:]
						u = graph[u][0]
					}
					paths = append(paths, nonBranchingPath)
				}
			}
		}
	}

	// Find isolated cycles
	for v := range graph {
		if !visited[v] {
			cycle := v
			u := v
			for {
				visited[u] = true
				if len(graph[u]) == 0 {
					break
				}
				w := graph[u][0]
				cycle += w[len(w)-1:]
				u = w
				if u == v {
					paths = append(paths, cycle)
					break
				}
			}
		}
	}

	return paths
}

func main() {
	file, err := os.Open("rosalind_ba3k.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read k-mers
	var patterns []string
	for scanner.Scan() {
		patterns = append(patterns, scanner.Text())
	}

	// Construct de Bruijn graph
	graph := constructDeBruijnGraph(patterns)

	// Find maximal non-branching paths
	contigs := findMaximalNonBranchingPaths(graph)

	// Sort contigs lexicographically
	sort.Strings(contigs)

	// Print result
	fmt.Println(strings.Join(contigs, " "))
}
