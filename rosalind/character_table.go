package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Node struct {
	name     string
	children []*Node
}

func parseNewick(s string) (*Node, error) {
	s = strings.TrimSpace(s)
	if !strings.HasSuffix(s, ";") {
		return nil, fmt.Errorf("invalid Newick format: missing semicolon at the end")
	}
	s = strings.TrimSuffix(s, ";")

	root := &Node{}
	current := root
	stack := []*Node{}
	name := ""

	for _, ch := range s {
		switch ch {
		case '(':
			if name != "" {
				current.name = name
				name = ""
			}
			newNode := &Node{}
			current.children = append(current.children, newNode)
			stack = append(stack, current)
			current = newNode
		case ')':
			if name != "" {
				current.children = append(current.children, &Node{name: name})
				name = ""
			}
			if len(stack) == 0 {
				return nil, fmt.Errorf("unbalanced parentheses")
			}
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case ',':
			if name != "" {
				current.children = append(current.children, &Node{name: name})
				name = ""
			}
		default:
			name += string(ch)
		}
	}

	if name != "" {
		if len(current.children) == 0 {
			current.name = name
		} else {
			current.children = append(current.children, &Node{name: name})
		}
	}

	if len(stack) != 0 {
		return nil, fmt.Errorf("unbalanced parentheses")
	}

	return root, nil
}

func collectTaxa(node *Node, taxa *[]string) {
	if len(node.children) == 0 {
		*taxa = append(*taxa, node.name)
	}
	for _, child := range node.children {
		collectTaxa(child, taxa)
	}
}

func generateCharacterTable(node *Node, taxa []string) [][]int {
	taxaMap := make(map[string]int)
	for i, taxon := range taxa {
		taxaMap[taxon] = i
	}

	var table [][]int
	var dfs func(*Node) []int

	dfs = func(n *Node) []int {
		if len(n.children) == 0 {
			row := make([]int, len(taxa))
			row[taxaMap[n.name]] = 1
			return row
		}

		row := make([]int, len(taxa))
		for _, child := range n.children {
			childRow := dfs(child)
			for i := range row {
				row[i] |= childRow[i]
			}
		}

		if !isTrival(row) {
			table = append(table, row)
		}

		return row
	}

	dfs(node)
	return table
}

func isTrival(row []int) bool {
	sum := 0
	for _, v := range row {
		sum += v
	}
	return sum <= 1 || sum >= len(row)-1
}

func main() {
	content, err := os.ReadFile("rosalind_ctbl.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	newickStr := strings.TrimSpace(string(content))

	root, err := parseNewick(newickStr)
	if err != nil {
		fmt.Println("Error parsing Newick string:", err)
		return
	}

	var taxa []string
	collectTaxa(root, &taxa)
	sort.Strings(taxa)
	fmt.Println(taxa)

	table := generateCharacterTable(root, taxa)
	fmt.Println(len(table))

	for _, row := range table {
		for _, v := range row {
			fmt.Print(v)
		}
		fmt.Println()
	}
}
