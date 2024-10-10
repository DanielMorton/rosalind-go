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

type Graph map[int][]Edge

func buildGraph(edges []string) Graph {
	graph := make(Graph)
	for _, edge := range edges {
		parts := strings.Split(edge, "->")
		from, _ := strconv.Atoi(parts[0])
		toWeight := strings.Split(parts[1], ":")
		to, _ := strconv.Atoi(toWeight[0])
		weight, _ := strconv.Atoi(toWeight[1])
		graph[from] = append(graph[from], Edge{to, weight})
	}
	return graph
}

func dfs(graph Graph, start, end, current int, visited map[int]bool, distance int) int {
	if current == end {
		return distance
	}
	visited[current] = true
	for _, edge := range graph[current] {
		if !visited[edge.to] {
			if d := dfs(graph, start, end, edge.to, visited, distance+edge.weight); d != -1 {
				return d
			}
		}
	}
	visited[current] = false
	return -1
}

func calculateDistances(graph Graph, n int) [][]int {
	distances := make([][]int, n)
	for i := range distances {
		distances[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			visited := make(map[int]bool)
			distance := dfs(graph, i, j, i, visited, 0)
			distances[i][j] = distance
			distances[j][i] = distance
		}
	}
	return distances
}

func main() {
	file, err := os.Open("rosalind_ba7a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var edges []string
	for scanner.Scan() {
		edges = append(edges, scanner.Text())
	}

	graph := buildGraph(edges)
	distances := calculateDistances(graph, n)

	// Print the distance matrix
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(distances[i][j])
		}
		fmt.Println()
	}
}
