package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type rotation struct {
	index int
	text  string
}

func constructBWT(text string) string {
	n := len(text)
	rotations := make([]rotation, n)

	// Create all cyclic rotations
	for i := 0; i < n; i++ {
		rotations[i] = rotation{
			index: i,
			text:  text[i:] + text[:i],
		}
	}

	// Sort rotations lexicographically
	sort.Slice(rotations, func(i, j int) bool {
		return rotations[i].text < rotations[j].text
	})

	// Construct BWT by taking the last character of each sorted rotation
	bwt := make([]byte, n)
	for i, r := range rotations {
		bwt[i] = text[(r.index-1+n)%n]
	}

	return string(bwt)
}

func main() {
	// Read input from file
	content, err := ioutil.ReadFile("rosalind_ba9i.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Trim whitespace and newlines
	text := strings.TrimSpace(string(content))

	// Construct the Burrows-Wheeler Transform
	bwt := constructBWT(text)

	// Print the result
	fmt.Println(bwt)
}
