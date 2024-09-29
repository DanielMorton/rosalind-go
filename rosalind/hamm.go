package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function to compute the Hamming distance between two DNA strings
func hammingDistance(s, t string) int {
	if len(s) != len(t) {
		panic("Strings must be of equal length")
	}

	distance := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			distance++
		}
	}
	return distance
}

func main() {
	// Open the file "rosalind_hamm.txt"
	file, err := os.Open("rosalind_hamm.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	var s, t string
	if scanner.Scan() {
		s = scanner.Text()
	}
	if scanner.Scan() {
		t = scanner.Text()
	}

	// Compute the Hamming distance
	distance := hammingDistance(s, t)

	// Print the result
	fmt.Println(distance)
}
