package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph map[int][]int

func readGraph(filename string) Graph {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	graph := make(Graph)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		node, _ := strconv.Atoi(parts[0])
		neighbors := strings.Split(parts[1], ",")
		for _, neighbor := range neighbors {
			neighborNode, _ := strconv.Atoi(strings.TrimSpace(neighbor))
			graph[node] = append(graph[node], neighborNode)
		}
	}
	return graph
}

func findStartEndNodes(graph Graph) (int, int) {
	inDegree := make(map[int]int)
	outDegree := make(map[int]int)

	for node, neighbors := range graph {
		outDegree[node] += len(neighbors)
		for _, neighbor := range neighbors {
			inDegree[neighbor]++
		}
	}

	var start, end int
	for node := range graph {
		if outDegree[node]-inDegree[node] == 1 {
			start = node
		} else if inDegree[node]-outDegree[node] == 1 {
			end = node
		}
	}

	return start, end
}

func findEulerianPath(graph Graph) []int {
	start, _ := findStartEndNodes(graph)

	stack := []int{start}
	path := []int{}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		if len(graph[current]) > 0 {
			next := graph[current][0]
			stack = append(stack, next)
			graph[current] = graph[current][1:]
		} else {
			path = append([]int{current}, path...)
			stack = stack[:len(stack)-1]
		}
	}

	return path
}

func main() {
	graph := readGraph("rosalind_ba3g.txt")
	path := findEulerianPath(graph)

	// Print the path
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(path)), "->"), "[]"))
}
