package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var aminoAcidMasses = []int{
	57, 71, 87, 97, 99, 101, 103, 113, 114, 115, 128, 129, 131, 137, 147, 156, 163, 186,
}

func countPeptidesWithMass(targetMass int) int64 {
	// Initialize the dynamic programming array
	dp := make([]int64, targetMass+1)
	dp[0] = 1 // Base case: one way to make mass 0 (empty peptide)

	// Fill the dp array
	for mass := 1; mass <= targetMass; mass++ {
		for _, aminoMass := range aminoAcidMasses {
			if mass >= aminoMass {
				dp[mass] += dp[mass-aminoMass]
			}
		}
	}

	return dp[targetMass]
}

func main() {
	file, err := os.Open("rosalind_ba4d.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	targetMass, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	count := countPeptidesWithMass(targetMass)
	fmt.Println(count)
}
