package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TrieNode struct {
	children map[byte]*TrieNode
	isLeaf   bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		isLeaf:   false,
	}
}

func constructTrie(patterns []string) *TrieNode {
	root := NewTrieNode()
	for _, pattern := range patterns {
		currentNode := root
		for i := 0; i < len(pattern); i++ {
			currentSymbol := pattern[i]
			if _, exists := currentNode.children[currentSymbol]; !exists {
				currentNode.children[currentSymbol] = NewTrieNode()
			}
			currentNode = currentNode.children[currentSymbol]
		}
		currentNode.isLeaf = true
	}
	return root
}

func prefixTrieMatching(text string, trie *TrieNode) bool {
	v := trie
	for i := 0; i < len(text); i++ {
		symbol := text[i]
		if v.isLeaf {
			return true
		}
		if child, exists := v.children[symbol]; exists {
			v = child
		} else {
			return false
		}
	}
	return v.isLeaf
}

func trieMatching(text string, trie *TrieNode) []int {
	positions := []int{}
	for i := 0; i < len(text); i++ {
		if prefixTrieMatching(text[i:], trie) {
			positions = append(positions, i)
		}
	}
	return positions
}

func main() {
	file, err := os.Open("rosalind_ba9b.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()

	var patterns []string
	for scanner.Scan() {
		patterns = append(patterns, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	trie := constructTrie(patterns)
	positions := trieMatching(text, trie)

	fmt.Println(strings.Trim(fmt.Sprint(positions), "[]"))
}
