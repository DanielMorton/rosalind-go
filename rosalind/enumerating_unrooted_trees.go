package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Left, Right *Node
	Value       string
}

func generateTrees(species []string) []*Node {
	if len(species) == 1 {
		return []*Node{{Value: species[0]}}
	}

	var trees []*Node
	remainingSpecies := append([]string{}, species[:len(species)-1]...)
	subTrees := generateTrees(remainingSpecies)
	for _, subTree := range subTrees {
		trees = append(trees, insertNode(subTree, species[len(species)-1])...)
	}
	return trees
}

func insertNode(tree *Node, value string) []*Node {
	var results []*Node

	// Insert as new root
	results = append(results, &Node{
		Left:  &Node{Value: value},
		Right: tree,
	})

	// If tree is not a leaf, recursively insert into left and right subtrees
	if tree.Left != nil && tree.Right != nil {
		leftInsertions := insertNode(tree.Left, value)
		for _, left := range leftInsertions {
			results = append(results, &Node{
				Left:  left,
				Right: tree.Right,
			})
		}

		rightInsertions := insertNode(tree.Right, value)
		for _, right := range rightInsertions {
			results = append(results, &Node{
				Left:  tree.Left,
				Right: right,
			})
		}
	}

	return results
}

func treeToNewick(node *Node) string {
	if node.Left == nil && node.Right == nil {
		return node.Value
	}
	left := treeToNewick(node.Left)
	right := treeToNewick(node.Right)
	return fmt.Sprintf("(%s,%s)", left, right)
}

func main() {
	file, err := os.Open("rosalind_eubt.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	species := strings.Fields(line)

	if len(species) < 3 {
		fmt.Println("Not enough species to form an unrooted binary tree")
		return
	}

	// Fix the last species as the root
	root := species[len(species)-1]
	species = species[:len(species)-1]

	trees := generateTrees(species)

	for _, tree := range trees {
		newick := fmt.Sprintf("(%s)%s;", treeToNewick(tree), root)
		fmt.Println(newick)
	}
}
