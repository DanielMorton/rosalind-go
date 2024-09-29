package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node struct {
	children []*Node
	name     string
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
			if index, ok := taxaMap[n.name]; ok {
				row[index] = 1
			}
			return row
		}

		childRows := make([][]int, len(n.children))
		for i, child := range n.children {
			childRows[i] = dfs(child)
		}

		row := make([]int, len(taxa))
		for _, childRow := range childRows {
			for i := range row {
				row[i] |= childRow[i]
			}
		}

		for _, childRow := range childRows {
			if !isTrival(childRow) {
				table = append(table, childRow)
			}
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

func countSharedSplits(table1, table2 [][]int) int {
	shared := 0
	for _, row1 := range table1 {
		for _, row2 := range table2 {
			if equalOrComplement(row1, row2) {
				shared++
				break
			}
		}
	}
	return shared
}

func equalOrComplement(row1, row2 []int) bool {
	if len(row1) != len(row2) {
		return false
	}
	equal := true
	complement := true
	for i := range row1 {
		if row1[i] != row2[i] {
			equal = false
		}
		if row1[i] == row2[i] {
			complement = false
		}
		if !equal && !complement {
			return false
		}
	}
	return true
}

func readLine(reader *bufio.Reader) (string, error) {
	var line strings.Builder
	for {
		part, isPrefix, err := reader.ReadLine()
		if err != nil {
			return "", err
		}
		line.Write(part)
		if !isPrefix {
			break
		}
	}
	return line.String(), nil
}

func main() {
	file, err := os.Open("rosalind_sptd.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// Read taxa
	taxaLine, err := readLine(reader)
	if err != nil {
		fmt.Println("Error reading taxa:", err)
		return
	}
	taxa := strings.Fields(taxaLine)

	// Read first tree
	tree1Line, err := readLine(reader)
	if err != nil {
		fmt.Println("Error reading first tree:", err)
		return
	}
	tree1, err := parseNewick(tree1Line)
	if err != nil {
		fmt.Println("Error parsing Newick string:", err)
		return
	}

	// Read second tree
	tree2Line, err := readLine(reader)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Error reading second tree:", err)
			return
		}
	}
	tree2, err := parseNewick(tree2Line)
	if err != nil {
		fmt.Println("Error parsing Newick string:", err)
		return
	}

	table1 := generateCharacterTable(tree1, taxa)
	table2 := generateCharacterTable(tree2, taxa)

	// Print tables for debugging
	/*fmt.Println("Table 1:")
	for _, row := range table1 {
		fmt.Println(row)
	}
	fmt.Println("Table 2:")
	for _, row := range table2 {
		fmt.Println(row)
	}*/

	sharedSplits := countSharedSplits(table1, table2)
	splitDistance := len(table1) + len(table2) - 2*sharedSplits

	fmt.Println("Split Distance:", splitDistance)
}
