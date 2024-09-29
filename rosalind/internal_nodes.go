package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// Read the input file
	data, err := ioutil.ReadFile("rosalind_inod.txt")
	if err != nil {
		panic(err)
	}

	// Parse the number of leaves (n)
	n, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		panic(err)
	}

	// Calculate the number of internal nodes
	internalNodes := n - 2

	// Output the result
	fmt.Println(internalNodes)
}
