package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Graph map[string][]string

func constructDeBruijnGraph(k int) Graph {
	graph := make(Graph)
	n := int(math.Pow(2, float64(k-1)))

	for i := 0; i < n; i++ {
		node := fmt.Sprintf("%0*b", k-1, i)
		graph[node] = []string{node[1:] + "0", node[1:] + "1"}
	}

	return graph
}

func findEulerianCycle(graph Graph) []string {
	// Start from any node
	start := ""
	for node := range graph {
		start = node
		break
	}

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

func reconstructCircularString(path []string, k int) string {
	if len(path) == 0 {
		return ""
	}
	result := path[0]
	for i := 1; i < len(path); i++ {
		result += string(path[i][len(path[i])-1])
	}
	return result[:len(result)-(k-1)]
}

func main() {
	// Read input
	content, err := ioutil.ReadFile("rosalind_ba3i.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	k, err := strconv.Atoi(strings.TrimSpace(string(content)))
	if err != nil {
		fmt.Println("Error parsing k:", err)
		os.Exit(1)
	}

	// Construct de Bruijn graph
	graph := constructDeBruijnGraph(k)

	// Find Eulerian cycle
	path := findEulerianCycle(graph)

	// Reconstruct circular string
	result := reconstructCircularString(path, k)

	// Print result
	fmt.Println(result)
}
