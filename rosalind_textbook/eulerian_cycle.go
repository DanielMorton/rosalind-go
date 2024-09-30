package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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

func findEulerianCycle(graph Graph) []int {
	// Choose a starting node
	start := -1
	for node := range graph {
		start = node
		break
	}

	stack := []int{start}
	path := []int{}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		if len(graph[current]) > 0 {
			// Choose the next node
			next := graph[current][0]
			stack = append(stack, next)
			// Remove the edge
			graph[current] = graph[current][1:]
		} else {
			// Add to path and backtrack
			path = append([]int{current}, path...)
			stack = stack[:len(stack)-1]
		}
	}

	return path
}

func main() {
	rand.Seed(time.Now().UnixNano())

	graph := readGraph("rosalind_ba3f.txt")
	cycle := findEulerianCycle(graph)

	// Print the cycle
	for i, node := range cycle {
		if i > 0 {
			fmt.Print("->")
		}
		fmt.Print(node)
	}
	fmt.Println()
}
