package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minCoins(money int, coins []int) int {
	// Create a slice to store the minimum number of coins for each amount
	dp := make([]int, money+1)

	// Initialize the slice with a value larger than any possible solution
	for i := range dp {
		dp[i] = money + 1
	}

	// Base case: 0 coins needed to make 0 change
	dp[0] = 0

	// Iterate through all amounts from 1 to money
	for i := 1; i <= money; i++ {
		// Try each coin
		for _, coin := range coins {
			if coin <= i {
				// Update dp[i] if using this coin leads to a smaller solution
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// If dp[money] is still money+1, it means no solution was found
	if dp[money] > money {
		return -1
	}

	return dp[money]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Read input from file
	file, err := os.Open("rosalind_ba5a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read money amount
	scanner.Scan()
	money, _ := strconv.Atoi(scanner.Text())

	// Read coin denominations
	scanner.Scan()
	coinStrs := strings.Split(scanner.Text(), ",")
	coins := make([]int, len(coinStrs))
	for i, s := range coinStrs {
		coins[i], _ = strconv.Atoi(s)
	}

	// Calculate and print the result
	result := minCoins(money, coins)
	fmt.Println(result)
}
