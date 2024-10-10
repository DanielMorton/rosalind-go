package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Edge struct {
	start, end int
	node       *Node
}

type Node struct {
	edges map[byte]*Edge
	depth int
}

type SuffixTree struct {
	text string
	root *Node
}

func newNode(depth int) *Node {
	return &Node{edges: make(map[byte]*Edge), depth: depth}
}

func newSuffixTree(text string) *SuffixTree {
	text += "$"
	st := &SuffixTree{text: text, root: newNode(0)}
	for i := 0; i < len(text); i++ {
		st.addSuffix(i)
	}
	return st
}

func (st *SuffixTree) addSuffix(start int) {
	cur := st.root
	for i := start; i < len(st.text); {
		edge, exists := cur.edges[st.text[i]]
		if !exists {
			cur.edges[st.text[i]] = &Edge{start: i, end: len(st.text) - 1, node: newNode(len(st.text) - i)}
			return
		}
		j := edge.start
		for j <= edge.end && i < len(st.text) && st.text[i] == st.text[j] {
			i++
			j++
		}
		if j > edge.end {
			cur = edge.node
			continue
		}
		split := newNode(cur.depth + j - edge.start)
		split.edges[st.text[j]] = &Edge{start: j, end: edge.end, node: edge.node}
		cur.edges[st.text[edge.start]] = &Edge{start: edge.start, end: j - 1, node: split}
		split.edges[st.text[i]] = &Edge{start: i, end: len(st.text) - 1, node: newNode(len(st.text) - i)}
		return
	}
}

func (st *SuffixTree) findLongestRepeat() string {
	maxLen := 0
	var longestRepeat string

	var dfs func(*Node, string)
	dfs = func(node *Node, currentString string) {
		if len(node.edges) > 1 && len(currentString) > maxLen {
			maxLen = len(currentString)
			longestRepeat = currentString
		}
		for _, edge := range node.edges {
			if edge.end-edge.start+1 > 0 { // Skip the terminal '$' edge
				dfs(edge.node, currentString+st.text[edge.start:edge.end+1])
			}
		}
	}

	dfs(st.root, "")
	return longestRepeat
}

func main() {
	content, err := ioutil.ReadFile("rosalind_ba9d.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := strings.TrimSpace(string(content))
	suffixTree := newSuffixTree(text)
	longestRepeat := suffixTree.findLongestRepeat()
	fmt.Println(longestRepeat)
}
