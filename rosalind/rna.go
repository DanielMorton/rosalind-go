package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("rosalind_rna.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	var dna string
	for scanner.Scan() {
		dna += scanner.Text()
	}

	// Check for any error during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Replace 'T' with 'U' to get the RNA string
	rna := strings.ReplaceAll(dna, "T", "U")

	// Output the RNA string
	fmt.Println(rna)
}
