package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Genotype    string
	Left, Right *Node
}

func parseNewick(s string) *Node {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, ";")
	return parseNode(s)
}

func parseNode(s string) *Node {
	if len(s) == 0 {
		return nil
	}
	if s[0] != '(' {
		return &Node{Genotype: s}
	}
	count := 0
	for i, c := range s {
		if c == '(' {
			count++
		} else if c == ')' {
			count--
		} else if c == ',' && count == 1 {
			left := parseNode(s[1:i])
			right := parseNode(s[i+1 : len(s)-1])
			return &Node{Left: left, Right: right}
		}
	}
	return nil
}

func calculateProbabilities(node *Node) (float64, float64, float64) {
	if node == nil {
		return 0, 0, 0
	}
	if node.Genotype != "" {
		switch node.Genotype {
		case "AA":
			return 1, 0, 0
		case "Aa":
			return 0.25, 0.5, 0.25
		case "aa":
			return 0, 0, 1
		}
	}

	lAA, lAa, laa := calculateProbabilities(node.Left)
	rAA, rAa, raa := calculateProbabilities(node.Right)

	pA_left := lAA + 0.5*lAa
	pa_left := laa + 0.5*lAa
	pA_right := rAA + 0.5*rAa
	pa_right := raa + 0.5*rAa

	pAA := pA_left * pA_right
	paa := pa_left * pa_right
	pAa := 1 - pAA - paa

	return pAA, pAa, paa
}

func main() {
	content, err := os.ReadFile("rosalind_mend.txt")
	if err != nil {
		panic(err)
	}

	// Convert content to string (Newick format tree)
	newick := string(content)

	root := parseNewick(newick)
	pAA, pAa, paa := calculateProbabilities(root)

	fmt.Printf("%.3f %.3f %.3f\n", pAA, pAa, paa)
}
