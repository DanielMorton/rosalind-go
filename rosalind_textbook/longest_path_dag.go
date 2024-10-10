package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	to     int
	weight int
}

func longestPath(graph map[int][]Edge, source, sink int) (int, []int) {
	// Initialize distances and predecessors
	dist := make(map[int]int)
	pred := make(map[int]int)
	for node := range graph {
		dist[node] = -1e9 // Initialize with a very small value
	}
	dist[source] = 0

	// Topological sort
	visited := make(map[int]bool)
	var order []int
	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true
		for _, edge := range graph[node] {
			if !visited[edge.to] {
				dfs(edge.to)
			}
		}
		order = append([]int{node}, order...)
	}
	dfs(source)

	// Process nodes in topological order
	for _, node := range order {
		for _, edge := range graph[node] {
			if dist[node]+edge.weight > dist[edge.to] {
				dist[edge.to] = dist[node] + edge.weight
				pred[edge.to] = node
			}
		}
	}

	// Reconstruct the path
	path := []int{sink}
	for node := sink; node != source; node = pred[node] {
		path = append([]int{pred[node]}, path...)
	}

	return dist[sink], path
}

func main() {
	// Read input from file
	file, err := os.Open("rosalind_ba5d.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read source and sink
	scanner.Scan()
	source, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	sink, _ := strconv.Atoi(scanner.Text())

	// Read graph
	graph := make(map[int][]Edge)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "->")
		from, _ := strconv.Atoi(parts[0])
		toWeight := strings.Split(parts[1], ":")
		to, _ := strconv.Atoi(toWeight[0])
		weight, _ := strconv.Atoi(toWeight[1])
		graph[from] = append(graph[from], Edge{to, weight})
	}

	// Calculate longest path
	length, path := longestPath(graph, source, sink)

	// Print result
	fmt.Println(length)
	for i, node := range path {
		if i > 0 {
			fmt.Print("->")
		}
		fmt.Print(node)
	}
	fmt.Println()
}
