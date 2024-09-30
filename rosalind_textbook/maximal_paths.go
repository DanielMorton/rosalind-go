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

func inDegree(graph Graph, node int) int {
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

func outDegree(graph Graph, node int) int {
	return len(graph[node])
}

func isOneInOneOut(graph Graph, node int) bool {
	return inDegree(graph, node) == 1 && outDegree(graph, node) == 1
}

func maximalNonBranchingPaths(graph Graph) [][]int {
	paths := [][]int{}
	visited := make(map[int]bool)

	for v := range graph {
		if !isOneInOneOut(graph, v) && outDegree(graph, v) > 0 {
			for _, w := range graph[v] {
				path := []int{v, w}
				visited[v] = true
				for isOneInOneOut(graph, w) {
					visited[w] = true
					u := graph[w][0]
					path = append(path, u)
					w = u
				}
				paths = append(paths, path)
			}
		}
	}

	// Find isolated cycles
	for v := range graph {
		if !visited[v] && isOneInOneOut(graph, v) {
			cycle := []int{v}
			w := v
			for {
				visited[w] = true
				u := graph[w][0]
				cycle = append(cycle, u)
				w = u
				if w == v {
					paths = append(paths, cycle)
					break
				}
			}
		}
	}

	return paths
}

func main() {
	file, err := os.Open("rosalind_ba3m.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := make(Graph)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		node, _ := strconv.Atoi(parts[0])
		neighbors := strings.Split(parts[1], ",")
		for _, neighbor := range neighbors {
			neighborNode, _ := strconv.Atoi(strings.TrimSpace(neighbor))
			graph[node] = append(graph[node], neighborNode)
		}
	}

	paths := maximalNonBranchingPaths(graph)

	// Sort paths for consistent output
	sort.Slice(paths, func(i, j int) bool {
		if paths[i][0] != paths[j][0] {
			return paths[i][0] < paths[j][0]
		}
		return paths[i][1] < paths[j][1]
	})

	// Print result
	for _, path := range paths {
		fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(path)), " -> "), "[]"))
	}
}
