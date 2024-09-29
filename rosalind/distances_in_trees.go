package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	name     string
	children []*Node
	parent   *Node
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
		node.name = s
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
				if i+1 < len(s) {
					node.name = s[i+1:]
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

func distance(a, b *Node) int {
	if a == b {
		return 0
	}
	aAncestors := make(map[*Node]int)
	dist := 0
	for node := a; node != nil; node = node.parent {
		aAncestors[node] = dist
		dist++
	}
	dist = 0
	for node := b; node != nil; node = node.parent {
		if aDistance, ok := aAncestors[node]; ok {
			return aDistance + dist
		}
		dist++
	}
	return -1 // Should never reach here if a and b are in the same tree
}

func main() {
	file, err := os.Open("rosalind_nwck.txt")
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

		dist := distance(nodeA, nodeB)
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
