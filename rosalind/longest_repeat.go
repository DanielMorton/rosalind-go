package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Edge structure representing the edges in the suffix tree
type Edge struct {
	parent string
	child  string
	start  int
	length int
}

// Read and parse input file
func parseInput(filename string) (string, int, []Edge) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	// First line contains the string and k
	s := lines[0]
	k, _ := strconv.Atoi(lines[1])

	// Parse the edges
	var edges []Edge
	for _, line := range lines[2:] {
		parts := strings.Fields(line)
		parent := parts[0]
		child := parts[1]
		start, _ := strconv.Atoi(parts[2])
		length, _ := strconv.Atoi(parts[3])
		edges = append(edges, Edge{parent, child, start, length})
	}

	return s, k, edges
}

// Depth-First Search to find the longest repeated substring
func dfs(node string, tree map[string][]Edge, s string, k int, substring string, leafCounts map[string]int, maxSubstring *string) int {
	if len(tree[node]) == 0 {
		// Leaf node
		leafCounts[node] = 1
		return 1
	}

	leafCount := 0
	for _, edge := range tree[node] {
		// Get the substring of the edge
		substr := s[edge.start-1 : edge.start-1+edge.length]

		// Recursively calculate leaf counts for the child
		childLeafCount := dfs(edge.child, tree, s, k, substring+substr, leafCounts, maxSubstring)
		leafCount += childLeafCount
	}

	leafCounts[node] = leafCount

	// Check if this node represents a valid repeated substring
	if leafCount >= k && len(substring) > len(*maxSubstring) {
		*maxSubstring = substring
	}

	return leafCount
}

func main() {
	// Read the input
	s, k, edges := parseInput("rosalind_lrep.txt")

	// Build the suffix tree from edges
	tree := make(map[string][]Edge)
	for _, edge := range edges {
		tree[edge.parent] = append(tree[edge.parent], edge)
	}

	// DFS to find the longest repeated substring
	leafCounts := make(map[string]int)
	maxSubstring := ""
	dfs("node1", tree, s, k, "", leafCounts, &maxSubstring)

	// Output the result
	fmt.Println(maxSubstring)
}
