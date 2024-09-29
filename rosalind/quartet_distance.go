package main

import (
	"bufio"
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

func countQuartets(row []int) int {
	zeros, ones := 0, 0
	for _, v := range row {
		if v == 0 {
			zeros++
		} else {
			ones++
		}
	}
	return (zeros * (zeros - 1) * ones * (ones - 1)) / 4
}

func rowToString(row []int) string {
	var sb strings.Builder
	for _, v := range row {
		sb.WriteString(fmt.Sprintf("%d", v))
	}
	return sb.String()
}

func totalQuartets(n int) int {
	return (n * (n - 1) * (n - 2) * (n - 3)) / 24
}

func main() {
	file, err := os.Open("rosalind_qrtd.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Skip the taxa line

	scanner.Scan()
	tree1, _ := parseNewick(scanner.Text())
	scanner.Scan()
	tree2, _ := parseNewick(scanner.Text())

	var taxa []string
	collectTaxa(tree1, &taxa)
	sort.Strings(taxa)

	table1 := generateCharacterTable(tree1, taxa)
	table2 := generateCharacterTable(tree2, taxa)

	// Convert tables to sets for efficient intersection
	set1 := make(map[string][]int)
	set2 := make(map[string][]int)

	for _, row := range table1 {
		set1[rowToString(row)] = row
	}
	for _, row := range table2 {
		set2[rowToString(row)] = row
	}

	// Count shared quartets
	sharedQuartets := 0
	for key, row := range set1 {
		if _, exists := set2[key]; exists {
			sharedQuartets += countQuartets(row)
		}
	}
	fmt.Println(sharedQuartets)

	// Calculate total quartets using the simple formula
	totalQuartets := totalQuartets(len(taxa))
	fmt.Println(totalQuartets)

	distance := 2*totalQuartets - 2*sharedQuartets

	fmt.Println(distance)
}
