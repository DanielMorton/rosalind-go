package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func factorial(n int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func calculateMaxMatchings(a, u, c, g int64) *big.Int {
	auPairs := new(big.Int).Div(
		factorial(max(a, u)),
		factorial(abs(a-u)),
	)
	cgPairs := new(big.Int).Div(
		factorial(max(c, g)),
		factorial(abs(c-g)),
	)
	return new(big.Int).Mul(auPairs, cgPairs)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	file, err := os.Open("rosalind_mmch.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rna string

	// Skip the first line (FASTA header)
	scanner.Scan()

	// Read the RNA sequence
	for scanner.Scan() {
		rna += scanner.Text()
	}

	// Count nucleotides
	var a, u, c, g int64
	for _, nucleotide := range rna {
		switch nucleotide {
		case 'A':
			a++
		case 'U':
			u++
		case 'C':
			c++
		case 'G':
			g++
		}
	}

	result := calculateMaxMatchings(a, u, c, g)
	fmt.Println(result.String())
}
