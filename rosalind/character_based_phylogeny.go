package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Label    string
	Children []*Node
}

func main() {
	species, characters := readInput("rosalind_chbp.txt")
	//fmt.Printf("Species: %v\n", species)
	//fmt.Printf("Characters: %v\n", characters)

	tree := buildTree(species, characters)

	newick := toNewick(tree) + ";"
	//fmt.Println("Newick format:")
	fmt.Println(newick)
}

func readInput(filename string) ([]string, []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	species := strings.Fields(scanner.Text())

	var characters []string
	for scanner.Scan() {
		characters = append(characters, scanner.Text())
	}

	return species, characters
}

func buildTree(species []string, characters []string) *Node {
	if len(species) == 1 {
		return &Node{Label: species[0]}
	}

	root := &Node{}

	for _, char := range characters {
		leftSpecies, rightSpecies := splitSpecies(species, char)
		if len(leftSpecies) > 0 && len(rightSpecies) > 0 {
			leftChars, rightChars := splitCharacters(characters, leftSpecies, rightSpecies)
			root.Children = append(root.Children,
				buildTree(leftSpecies, leftChars),
				buildTree(rightSpecies, rightChars))
			break
		}
	}

	if len(root.Children) == 0 {
		for _, s := range species {
			root.Children = append(root.Children, &Node{Label: s})
		}
	}

	return root
}

func splitSpecies(species []string, character string) ([]string, []string) {
	var left, right []string
	for i, c := range character {
		if c == '0' {
			left = append(left, species[i])
		} else {
			right = append(right, species[i])
		}
	}
	return left, right
}

func splitCharacters(characters []string, leftSpecies, rightSpecies []string) ([]string, []string) {
	var leftChars, rightChars []string
	for _, char := range characters {
		leftChar, rightChar := "", ""
		for i, c := range char {
			if i < len(leftSpecies) {
				leftChar += string(c)
			} else {
				rightChar += string(c)
			}
		}
		leftChars = append(leftChars, leftChar)
		rightChars = append(rightChars, rightChar)
	}
	return leftChars, rightChars
}

func toNewick(node *Node) string {
	if node.Label != "" {
		return node.Label
	}

	childStrings := make([]string, len(node.Children))
	for i, child := range node.Children {
		childStrings[i] = toNewick(child)
	}

	return "(" + strings.Join(childStrings, ",") + ")"
}
