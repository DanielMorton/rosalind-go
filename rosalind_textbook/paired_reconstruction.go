package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Graph map[string][]string

func constructDeBruijnGraph(pairedReads []string) Graph {
	graph := make(Graph)
	for _, read := range pairedReads {
		parts := strings.Split(read, "|")
		prefix := parts[0][:len(parts[0])-1] + "|" + parts[1][:len(parts[1])-1]
		suffix := parts[0][1:] + "|" + parts[1][1:]
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

func reconstructString(path []string, k, d int) string {
	firstPart := ""
	secondPart := ""
	for i, node := range path {
		parts := strings.Split(node, "|")
		if i == 0 {
			firstPart = parts[0]
			secondPart = parts[1]
		} else {
			firstPart += string(parts[0][len(parts[0])-1])
			secondPart += string(parts[1][len(parts[1])-1])
		}
	}
	return firstPart + secondPart[len(secondPart)-k-d:]
}

func main() {
	file, err := os.Open("rosalind_ba3j.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read k and d
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	k := 0
	d := 0
	fmt.Sscanf(params[0], "%d", &k)
	fmt.Sscanf(params[1], "%d", &d)

	// Read paired k-mers
	var pairedReads []string
	for scanner.Scan() {
		pairedReads = append(pairedReads, scanner.Text())
	}

	// Construct de Bruijn graph
	graph := constructDeBruijnGraph(pairedReads)

	// Find Eulerian path
	path := findEulerianPath(graph)

	// Reconstruct string
	result := reconstructString(path, k, d)

	// Print result
	fmt.Println(result)
}
