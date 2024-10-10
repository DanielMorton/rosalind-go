package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Graph map[int][]int

func topologicalSort(graph Graph) []int {
	visited := make(map[int]bool)
	stack := []int{}

	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
		stack = append([]int{node}, stack...)
	}

	// Get all nodes in the graph
	nodes := make([]int, 0, len(graph))
	for node := range graph {
		nodes = append(nodes, node)
	}

	// Sort nodes to ensure deterministic output
	sort.Ints(nodes)

	// Perform DFS on each unvisited node
	for _, node := range nodes {
		if !visited[node] {
			dfs(node)
		}
	}

	return stack
}

func main() {
	file, err := os.Open("rosalind_ba5n.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := make(Graph)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			continue
		}

		from, _ := strconv.Atoi(parts[0])
		toNodes := strings.Split(parts[1], ",")
		for _, toStr := range toNodes {
			to, _ := strconv.Atoi(strings.TrimSpace(toStr))
			graph[from] = append(graph[from], to)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	order := topologicalSort(graph)

	// Print the result
	for i, node := range order {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(node)
	}
	fmt.Println()
}
