package main

import (
	"bufio"
	"fmt"
	"os"
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

type SuffixTree struct {
	root *Node
	text string
}

func newSuffixTree(text string) *SuffixTree {
	return &SuffixTree{
		root: newNode(-1, -1),
		text: text,
	}
}

func (st *SuffixTree) addSuffix(suffix int) {
	current := st.root
	i := suffix
	for i < len(st.text) {
		edge := st.text[i]
		if child, exists := current.children[edge]; exists {
			childStart := child.start
			childEnd := child.end
			j := 0
			for i < len(st.text) && childStart+j < childEnd && st.text[i] == st.text[childStart+j] {
				i++
				j++
			}
			if childStart+j == childEnd {
				current = child
				continue
			}
			splitNode := newNode(child.start, child.start+j)
			current.children[edge] = splitNode
			child.start += j
			splitNode.children[st.text[child.start]] = child
			if i < len(st.text) {
				splitNode.children[st.text[i]] = newNode(i, len(st.text))
			}
			return
		} else {
			current.children[edge] = newNode(i, len(st.text))
			return
		}
	}
}

func (st *SuffixTree) build() {
	for i := 0; i < len(st.text); i++ {
		st.addSuffix(i)
	}
}

func (st *SuffixTree) getEdgeLabels() []string {
	var labels []string
	var dfs func(*Node)
	dfs = func(node *Node) {
		for _, child := range node.children {
			if child.start != -1 {
				labels = append(labels, st.text[child.start:child.end])
			}
			dfs(child)
		}
	}
	dfs(st.root)
	return labels
}

func main() {
	file, err := os.Open("rosalind_ba9c.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	st := newSuffixTree(text)
	st.build()
	labels := st.getEdgeLabels()

	for _, label := range labels {
		fmt.Println(label)
	}
}
