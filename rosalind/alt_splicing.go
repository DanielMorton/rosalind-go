package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MOD = 1000000

func sumCombinations(n, m int) int {
	// Initialize the DP table
	dp := make([]int, n+1)
	dp[0] = 1

	// Fill the DP table
	for i := 1; i <= n; i++ {
		for j := i; j > 0; j-- {
			dp[j] = (dp[j] + dp[j-1]) % MOD
		}
	}

	// Sum the combinations from m to n
	sum := 0
	for i := m; i <= n; i++ {
		sum = (sum + dp[i]) % MOD
	}

	return sum
}

func main() {
	file, err := os.Open("rosalind_aspc.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	parts := strings.Fields(line)
	if len(parts) != 2 {
		fmt.Println("Invalid input format")
		return
	}

	n, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println("Error parsing n:", err)
		return
	}

	m, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Error parsing m:", err)
		return
	}

	result := sumCombinations(n, m)
	fmt.Println(result)
}
