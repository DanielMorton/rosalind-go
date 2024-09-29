package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
)

// Function to calculate factorial of a number
func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func main() {
	// Read input from file
	data, err := ioutil.ReadFile("rosalind_pmch.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the RNA sequence (ignore the label line starting with '>')
	lines := strings.Split(string(data), "\n")
	var rna string
	for _, line := range lines {
		if !strings.HasPrefix(line, ">") {
			rna += line
		}
	}

	// Count occurrences of 'A', 'U', 'C', 'G'
	countA := strings.Count(rna, "A")
	countU := strings.Count(rna, "U")
	countC := strings.Count(rna, "C")
	countG := strings.Count(rna, "G")

	// Check if the numbers of 'A' and 'U', 'C' and 'G' are equal
	if countA != countU || countC != countG {
		fmt.Println("No perfect matching possible")
		return
	}

	// Calculate the total number of perfect matchings
	totalMatchings := new(big.Int).Mul(factorial(countA), factorial(countC))

	// Print the result
	fmt.Println(totalMatchings)
}
