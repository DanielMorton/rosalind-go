package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Read input from file
	data, err := ioutil.ReadFile("rosalind_tree.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Split the input into lines
	lines := strings.Split(string(data), "\n")

	// The first line contains the number of nodes, n
	n := 0
	fmt.Sscanf(lines[0], "%d", &n)

	// The number of edges is simply the number of remaining non-empty lines
	edges := 0
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) != "" {
			edges++
		}
	}

	// The minimum number of edges to add to form a tree is (n - 1) - current number of edges
	fmt.Println((n - 1) - edges)
}
