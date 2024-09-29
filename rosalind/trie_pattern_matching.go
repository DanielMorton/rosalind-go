package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TrieNode represents a node in the trie
type TrieNode struct {
	Children map[rune]*TrieNode
	ID       int
}

// Trie represents the trie structure
type Trie struct {
	Root      *TrieNode
	NextID    int
	Adjacency []string
}

// NewTrie creates a new Trie instance
func NewTrie() *Trie {
	root := &TrieNode{Children: make(map[rune]*TrieNode)}
	return &Trie{Root: root, NextID: 2}
}

// Insert adds a new string to the trie
func (t *Trie) Insert(s string) {
	current := t.Root
	for _, char := range s {
		if _, exists := current.Children[char]; !exists {
			newNode := &TrieNode{Children: make(map[rune]*TrieNode)}
			current.Children[char] = newNode
		}
		current = current.Children[char]
	}
}

// BuildTrie builds the trie from a collection of patterns
func (t *Trie) BuildTrie(patterns []string) {
	for _, s := range patterns {
		t.Insert(s)
	}
}

// TraverseTrie traverses the trie to generate the adjacency list
func (t *Trie) TraverseTrie(node *TrieNode, parentID int) {
	for char, child := range node.Children {
		child.ID = t.NextID
		t.NextID++
		t.Adjacency = append(t.Adjacency, fmt.Sprintf("%d %d %c", parentID, child.ID, char))
		t.TraverseTrie(child, child.ID)
	}
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_trie.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var patterns []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 {
			patterns = append(patterns, trimmedLine)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Initialize and build the trie
	trie := NewTrie()
	trie.BuildTrie(patterns)

	// Traverse the trie to generate the adjacency list
	trie.TraverseTrie(trie.Root, 1)

	// Print the adjacency list
	for _, entry := range trie.Adjacency {
		fmt.Println(entry)
	}
}
