package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
)

func isPair(a, b byte) bool {
	return (a == 'A' && b == 'U') ||
		(a == 'U' && b == 'A') ||
		(a == 'C' && b == 'G') ||
		(a == 'G' && b == 'C') ||
		(a == 'U' && b == 'G') ||
		(a == 'G' && b == 'U')
}

func countRNASecondaryStructures(rna string) *big.Int {
	n := len(rna)
	dp := make([][]*big.Int, n)
	for i := range dp {
		dp[i] = make([]*big.Int, n)
		for j := range dp[i] {
			dp[i][j] = big.NewInt(0)
		}
	}

	// Initialize base cases
	for i := 0; i < n; i++ {
		dp[i][i].SetInt64(1)
		if i < n-1 {
			dp[i][i+1].SetInt64(1)
		}
	}

	// Fill the dp table
	for length := 2; length < n; length++ {
		for i := 0; i < n-length; i++ {
			j := i + length

			// Case 1: j is unpaired
			dp[i][j].Set(dp[i][j-1])

			// Case 2: j is paired with some k
			for k := i; k < j-3; k++ {
				if isPair(rna[k], rna[j]) {
					if k == i {
						dp[i][j].Add(dp[i][j], dp[i+1][j-1])
					} else {
						temp := new(big.Int).Mul(dp[i][k-1], dp[k+1][j-1])
						dp[i][j].Add(dp[i][j], temp)
					}
				}
			}
		}
	}

	return dp[0][n-1]
}

func main() {
	// Read input from file
	content, err := ioutil.ReadFile("rosalind_rnas.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Clean the input
	rna := strings.TrimSpace(string(content))

	// Calculate and print the result
	result := countRNASecondaryStructures(rna)
	fmt.Println(result.String())
}
