package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	label    int
	children []*Node
	age      float64
	size     int
}

type Edge struct {
	from, to int
	weight   float64
}

func UPGMA(D [][]float64, n int) *Node {
	clusters := make([]*Node, n)
	for i := 0; i < n; i++ {
		clusters[i] = &Node{label: i, age: 0, size: 1}
	}

	nodeCount := n
	for len(clusters) > 1 {
		// Find the two closest clusters
		minDist := math.Inf(1)
		var minI, minJ int
		for i := 0; i < len(clusters)-1; i++ {
			for j := i + 1; j < len(clusters); j++ {
				if D[i][j] < minDist {
					minDist = D[i][j]
					minI, minJ = i, j
				}
			}
		}

		// Create a new cluster
		newNode := &Node{
			label:    nodeCount,
			children: []*Node{clusters[minI], clusters[minJ]},
			age:      minDist / 2,
			size:     clusters[minI].size + clusters[minJ].size,
		}
		nodeCount++

		// Update distance matrix
		newD := make([][]float64, len(clusters)-1)
		for i := range newD {
			newD[i] = make([]float64, len(clusters)-1)
		}

		newIndex := 0
		for i := 0; i < len(clusters); i++ {
			if i == minI || i == minJ {
				continue
			}
			newJIndex := 0
			for j := 0; j < len(clusters); j++ {
				if j == minI || j == minJ {
					continue
				}
				newD[newIndex][newJIndex] = D[i][j]
				newJIndex++
			}
			// Calculate distance to new cluster
			newDist := (D[i][minI]*float64(clusters[minI].size) + D[i][minJ]*float64(clusters[minJ].size)) /
				float64(newNode.size)
			newD[newIndex][len(newD)-1] = newDist
			newD[len(newD)-1][newIndex] = newDist
			newIndex++
		}

		// Update clusters
		newClusters := make([]*Node, 0, len(clusters)-1)
		for i, cluster := range clusters {
			if i != minI && i != minJ {
				newClusters = append(newClusters, cluster)
			}
		}
		newClusters = append(newClusters, newNode)
		clusters = newClusters
		D = newD
	}

	return clusters[0]
}

func buildAdjacencyList(root *Node) []Edge {
	edges := []Edge{}
	var dfs func(*Node)
	dfs = func(node *Node) {
		for _, child := range node.children {
			edges = append(edges, Edge{
				from:   node.label,
				to:     child.label,
				weight: node.age - child.age,
			})
			edges = append(edges, Edge{
				from:   child.label,
				to:     node.label,
				weight: node.age - child.age,
			})
			dfs(child)
		}
	}
	dfs(root)
	return edges
}

func main() {
	file, err := os.Open("rosalind_ba7d.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	D := make([][]float64, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		row := strings.Fields(scanner.Text())
		D[i] = make([]float64, n)
		for j, val := range row {
			D[i][j], _ = strconv.ParseFloat(val, 64)
		}
	}

	root := UPGMA(D, n)
	edges := buildAdjacencyList(root)

	// Sort edges for consistent output
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].from == edges[j].from {
			return edges[i].to < edges[j].to
		}
		return edges[i].from < edges[j].from
	})

	for _, edge := range edges {
		fmt.Printf("%d->%d:%.3f\n", edge.from, edge.to, edge.weight)
	}
}
