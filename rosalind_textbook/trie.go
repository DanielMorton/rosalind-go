package main

import (
	"bufio"
	"fmt"
	"os"
)

type TrieNode struct {
	children map[byte]*TrieNode
	nodeID   int
}

func NewTrieNode(id int) *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		nodeID:   id,
	}
}

func constructTrie(patterns []string) *TrieNode {
	root := NewTrieNode(0)
	nextID := 1

	for _, pattern := range patterns {
		currentNode := root
		for i := 0; i < len(pattern); i++ {
			currentSymbol := pattern[i]
			if _, exists := currentNode.children[currentSymbol]; !exists {
				newNode := NewTrieNode(nextID)
				nextID++
				currentNode.children[currentSymbol] = newNode
			}
			currentNode = currentNode.children[currentSymbol]
		}
	}

	return root
}

func printTrieEdges(node *TrieNode, parentID int) {
	for symbol, child := range node.children {
		fmt.Printf("%d->%d:%c\n", parentID, child.nodeID, symbol)
		printTrieEdges(child, child.nodeID)
	}
}

func main() {
	// Read input from file
	file, err := os.Open("rosalind_ba9a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var patterns []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		patterns = append(patterns, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Construct the trie
	root := constructTrie(patterns)

	// Print the trie edges
	printTrieEdges(root, root.nodeID)
}
