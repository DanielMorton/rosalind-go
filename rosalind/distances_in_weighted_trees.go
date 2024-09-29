package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	name     string
	children []*Node
	parent   *Node
	weight   float64
}

func parseNewick(s string) *Node {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, ";")
	root, _ := parseNewickHelper(s)
	return root
}

func parseNewickHelper(s string) (*Node, string) {
	node := &Node{}
	if len(s) == 0 {
		return node, ""
	}
	if s[0] != '(' {
		parts := strings.SplitN(s, ":", 2)
		node.name = parts[0]
		if len(parts) > 1 {
			weight, _ := strconv.ParseFloat(parts[1], 64)
			node.weight = weight
		}
		return node, ""
	}
	level := 0
	start := 1
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			level++
		} else if s[i] == ')' {
			level--
			if level == -1 {
				child, _ := parseNewickHelper(s[start:i])
				child.parent = node
				node.children = append(node.children, child)
				remaining := s[i+1:]
				parts := strings.SplitN(remaining, ":", 2)
				if len(parts) > 1 {
					node.name = parts[0]
					weight, _ := strconv.ParseFloat(parts[1], 64)
					node.weight = weight
				}
				return node, ""
			}
		} else if s[i] == ',' && level == 0 {
			child, _ := parseNewickHelper(s[start:i])
			child.parent = node
			node.children = append(node.children, child)
			start = i + 1
		}
	}
	return node, ""
}

func findNode(root *Node, name string) *Node {
	if root.name == name {
		return root
	}
	for _, child := range root.children {
		if found := findNode(child, name); found != nil {
			return found
		}
	}
	return nil
}

func distance(a, b *Node) float64 {
	if a == b {
		return 0
	}
	aAncestors := make(map[*Node]float64)
	dist := 0.0
	for node := a; node != nil; node = node.parent {
		aAncestors[node] = dist
		dist += node.weight
	}
	dist = 0
	for node := b; node != nil; node = node.parent {
		if aDistance, ok := aAncestors[node]; ok {
			return aDistance + dist
		}
		dist += node.weight
	}
	return -1 // Should never reach here if a and b are in the same tree
}

func main() {
	file, err := os.Open("rosalind_nkew.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var results []int

	for scanner.Scan() {
		newickStr := scanner.Text()
		if newickStr == "" {
			continue
		}
		root := parseNewick(newickStr)

		scanner.Scan()
		nodePair := strings.Fields(scanner.Text())
		if len(nodePair) != 2 {
			fmt.Println("Invalid node pair")
			continue
		}

		nodeA := findNode(root, nodePair[0])
		nodeB := findNode(root, nodePair[1])
		if nodeA == nil || nodeB == nil {
			fmt.Printf("Node not found: %v or %v\n", nodePair[0], nodePair[1])
			continue
		}

		dist := int(distance(nodeA, nodeB))
		results = append(results, dist)

		// Skip the empty line
		scanner.Scan()
	}

	// Print results
	for i, result := range results {
		fmt.Print(result)
		if i < len(results)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
