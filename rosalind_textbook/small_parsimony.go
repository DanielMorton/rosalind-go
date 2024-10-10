package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Node struct {
	label    string
	children []*Node
	parent   *Node
	sk       [4]int
}

const ALPHABET = "ACGT"

func SmallParsimony(root *Node) int {
	if len(root.children) == 0 {
		for i, ch := range ALPHABET {
			if root.label[0] == byte(ch) {
				root.sk[i] = 0
			} else {
				root.sk[i] = math.MaxInt32
			}
		}
		return 0
	}

	score := 0
	for _, child := range root.children {
		score += SmallParsimony(child)
	}

	for i := range ALPHABET {
		min1 := math.MaxInt32
		min2 := math.MaxInt32
		for j := range ALPHABET {
			score1 := root.children[0].sk[j]
			score2 := root.children[1].sk[j]
			if i != j {
				score1++
				score2++
			}
			min1 = min(min1, score1)
			min2 = min(min2, score2)
		}
		root.sk[i] = min1 + min2
	}

	minScore := min(root.sk[0], root.sk[1], root.sk[2], root.sk[3])
	return score + minScore
}

func min(a int, b ...int) int {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}

func assignLabels(node *Node) {
	if len(node.children) == 0 {
		return
	}

	minIndex := 0
	for i := 1; i < 4; i++ {
		if node.sk[i] < node.sk[minIndex] {
			minIndex = i
		}
	}
	node.label = string(ALPHABET[minIndex])

	for _, child := range node.children {
		assignLabels(child)
	}
}

func buildTree(adjacencyList map[string][]string) *Node {
	nodes := make(map[string]*Node)

	var getNode func(string) *Node
	getNode = func(label string) *Node {
		if node, exists := nodes[label]; exists {
			return node
		}
		node := &Node{label: label}
		nodes[label] = node
		return node
	}

	var root *Node
	for parent, children := range adjacencyList {
		parentNode := getNode(parent)
		if root == nil {
			root = parentNode
		}
		for _, child := range children {
			childNode := getNode(child)
			parentNode.children = append(parentNode.children, childNode)
			childNode.parent = parentNode
		}
	}

	return root
}

func hammingDistance(s1, s2 string) int {
	dist := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			dist++
		}
	}
	return dist
}

func main() {
	file, err := os.Open("rosalind_ba7f.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Skip the first line

	adjacencyList := make(map[string][]string)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "->")
		if len(parts) == 2 {
			adjacencyList[parts[0]] = append(adjacencyList[parts[0]], parts[1])
		}
	}

	root := buildTree(adjacencyList)

	totalScore := 0
	dnaLength := len(root.children[0].label)
	for i := 0; i < dnaLength; i++ {
		for node := range adjacencyList {
			if len(adjacencyList[node]) == 0 {
				nodes[node].label = string(nodes[node].label[i])
			} else {
				nodes[node].label = ""
			}
		}
		totalScore += SmallParsimony(root)
		assignLabels(root)
		for node := range adjacencyList {
			if len(adjacencyList[node]) == 0 {
				nodes[node].label += string(nodes[node].label[0])
			}
		}
	}

	fmt.Println(totalScore)

	var printEdges func(*Node)
	printEdges = func(node *Node) {
		for _, child := range node.children {
			dist := hammingDistance(node.label, child.label)
			fmt.Printf("%s->%s:%d\n", node.label, child.label, dist)
			fmt.Printf("%s->%s:%d\n", child.label, node.label, dist)
			printEdges(child)
		}
	}

	printEdges(root)
}
