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
	children map[int]float64
}

func NeighborJoining(D [][]float64, n int) map[int]*Node {
	tree := make(map[int]*Node)
	for i := 0; i < n; i++ {
		tree[i] = &Node{label: i, children: make(map[int]float64)}
	}

	activeNodes := make([]int, n)
	for i := range activeNodes {
		activeNodes[i] = i
	}

	for len(activeNodes) > 2 {
		m := len(activeNodes)

		// Calculate total distances
		totalDist := make([]float64, m)
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				totalDist[i] += D[i][j]
			}
		}

		// Find minimum element in neighbor-joining matrix
		minValue := math.Inf(1)
		var minI, minJ int
		for i := 0; i < m-1; i++ {
			for j := i + 1; j < m; j++ {
				value := float64(m-2)*D[i][j] - totalDist[i] - totalDist[j]
				if value < minValue {
					minValue = value
					minI, minJ = i, j
				}
			}
		}

		// Calculate delta and limb lengths
		delta := (totalDist[minI] - totalDist[minJ]) / float64(m-2)
		limbLengthI := (D[minI][minJ] + delta) / 2
		limbLengthJ := (D[minI][minJ] - delta) / 2

		// Create new node
		newNode := &Node{label: len(tree), children: make(map[int]float64)}
		tree[newNode.label] = newNode

		// Connect new node to minI and minJ
		newNode.children[activeNodes[minI]] = limbLengthI
		newNode.children[activeNodes[minJ]] = limbLengthJ
		tree[activeNodes[minI]].children[newNode.label] = limbLengthI
		tree[activeNodes[minJ]].children[newNode.label] = limbLengthJ

		// Update distance matrix
		newD := make([][]float64, m-1)
		for i := range newD {
			newD[i] = make([]float64, m-1)
		}

		newIndex := 0
		for i := 0; i < m; i++ {
			if i == minI || i == minJ {
				continue
			}
			newJIndex := 0
			for j := 0; j < m; j++ {
				if j == minI || j == minJ {
					continue
				}
				newD[newIndex][newJIndex] = D[i][j]
				newJIndex++
			}
			newD[newIndex][m-2] = (D[i][minI] + D[i][minJ] - D[minI][minJ]) / 2
			newD[m-2][newIndex] = newD[newIndex][m-2]
			newIndex++
		}

		// Update active nodes
		newActiveNodes := make([]int, 0, m-1)
		for i, node := range activeNodes {
			if i != minI && i != minJ {
				newActiveNodes = append(newActiveNodes, node)
			}
		}
		newActiveNodes = append(newActiveNodes, newNode.label)

		D = newD
		activeNodes = newActiveNodes
	}

	// Connect the last two nodes
	if len(activeNodes) == 2 {
		dist := D[0][1]
		tree[activeNodes[0]].children[activeNodes[1]] = dist
		tree[activeNodes[1]].children[activeNodes[0]] = dist
	}

	return tree
}

func main() {
	file, err := os.Open("rosalind_ba7e.txt")
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

	tree := NeighborJoining(D, n)

	// Prepare output
	edges := make([]string, 0)
	for i, node := range tree {
		for j, dist := range node.children {
			if i < j {
				edges = append(edges, fmt.Sprintf("%d->%d:%.3f", i, j, dist))
				edges = append(edges, fmt.Sprintf("%d->%d:%.3f", j, i, dist))
			}
		}
	}

	// Sort edges for consistent output
	sort.Strings(edges)

	// Print output
	for _, edge := range edges {
		fmt.Println(edge)
	}
}
