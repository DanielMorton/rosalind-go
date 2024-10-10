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

func limb(D [][]int, j, n int) int {
	minLimbLength := 1 << 30
	for i := 0; i < n; i++ {
		if i == j {
			continue
		}
		for k := i + 1; k < n; k++ {
			if k == j {
				continue
			}
			length := (D[i][j] + D[j][k] - D[i][k]) / 2
			if length < minLimbLength {
				minLimbLength = length
			}
		}
	}
	return minLimbLength
}

func additivePhylogeny(D [][]int, n int, nodeCount *int) Graph {
	if n == 2 {
		graph := make(Graph)
		graph[0] = append(graph[0], Edge{1, D[0][1]})
		graph[1] = append(graph[1], Edge{0, D[0][1]})
		return graph
	}

	limbLength := limb(D, n-1, n)
	for j := 0; j < n-1; j++ {
		D[j][n-1] -= limbLength
		D[n-1][j] = D[j][n-1]
	}

	var i, k int
	x := 0
	for i = 0; i < n-1; i++ {
		for k = i + 1; k < n-1; k++ {
			if D[i][k] == D[i][n-1]+D[n-1][k] {
				x = D[i][n-1]
				goto found
			}
		}
	}
found:

	reducedD := make([][]int, n-1)
	for i := range reducedD {
		reducedD[i] = make([]int, n-1)
		copy(reducedD[i], D[i][:n-1])
	}

	T := additivePhylogeny(reducedD, n-1, nodeCount)

	path := findPath(T, i, k)
	var v int
	for j := 0; j < len(path)-1; j++ {
		if x <= T[path[j]][0].weight {
			v = *nodeCount
			*nodeCount++
			weight := T[path[j]][0].weight
			T[path[j]] = []Edge{{v, x}}
			T[v] = append(T[v], Edge{path[j], x})
			T[v] = append(T[v], Edge{path[j+1], weight - x})
			T[path[j+1]] = append(T[path[j+1]], Edge{v, weight - x})
			break
		}
		x -= T[path[j]][0].weight
	}

	if _, exists := T[v]; !exists {
		v = path[len(path)-1]
	}

	T[v] = append(T[v], Edge{n - 1, limbLength})
	T[n-1] = append(T[n-1], Edge{v, limbLength})

	return T
}

func findPath(graph Graph, start, end int) []int {
	visited := make(map[int]bool)
	path := []int{}
	if dfs(graph, start, end, visited, &path) {
		return path
	}
	return nil
}

func dfs(graph Graph, current, end int, visited map[int]bool, path *[]int) bool {
	visited[current] = true
	*path = append(*path, current)

	if current == end {
		return true
	}

	for _, edge := range graph[current] {
		if !visited[edge.to] {
			if dfs(graph, edge.to, end, visited, path) {
				return true
			}
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func main() {
	file, err := os.Open("rosalind_ba7c.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	D := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		row := strings.Fields(scanner.Text())
		D[i] = make([]int, n)
		for j, val := range row {
			D[i][j], _ = strconv.Atoi(val)
		}
	}

	nodeCount := n
	T := additivePhylogeny(D, n, &nodeCount)

	for i := 0; i < nodeCount; i++ {
		for _, edge := range T[i] {
			fmt.Printf("%d->%d:%d\n", i, edge.to, edge.weight)
		}
	}
}
