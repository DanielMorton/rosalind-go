package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Node struct {
	children map[byte]*Node
	start    int
	end      int
}

func newNode(start, end int) *Node {
	return &Node{
		children: make(map[byte]*Node),
		start:    start,
		end:      end,
	}
}

func buildSuffixTree(s string) *Node {
	root := newNode(-1, -1)
	for i := 0; i < len(s); i++ {
		addSuffix(root, s, i)
	}
	return root
}

func addSuffix(root *Node, s string, start int) {
	curr := root
	for i := start; i < len(s); {
		child, exists := curr.children[s[i]]
		if !exists {
			curr.children[s[i]] = newNode(i, len(s)-1)
			return
		}
		j := 0
		for i+j < len(s) && j < child.end-child.start+1 && s[i+j] == s[child.start+j] {
			j++
		}
		if j == child.end-child.start+1 {
			i += j
			curr = child
		} else {
			splitNode := newNode(child.start, child.start+j-1)
			curr.children[s[child.start]] = splitNode
			splitNode.children[s[child.start+j]] = child
			child.start += j
			splitNode.children[s[i+j]] = newNode(i+j, len(s)-1)
			return
		}
	}
}

type Edge struct {
	label string
	start int
	end   int
}

func collectEdges(node *Node, s string, edges *[]Edge) {
	for _, child := range node.children {
		*edges = append(*edges, Edge{
			label: s[child.start : child.end+1],
			start: child.start,
			end:   child.end,
		})
		collectEdges(child, s, edges)
	}
}

func main() {
	file, err := os.Open("rosalind_suff.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s := scanner.Text()

	root := buildSuffixTree(s)

	var edges []Edge
	collectEdges(root, s, &edges)

	// Sort edges first by label, then by start index
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].label == edges[j].label {
			return edges[i].start < edges[j].start
		}
		return edges[i].label < edges[j].label
	})

	// Print edges, keeping duplicates
	for _, edge := range edges {
		fmt.Println(edge.label)
	}
}
