package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("rosalind_lgis.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the input
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text()) // the length of the permutation
	scanner.Scan()
	permutation := make([]int, n)
	strs := strings.Fields(scanner.Text())
	for i := 0; i < n; i++ {
		permutation[i], _ = strconv.Atoi(strs[i])
	}

	// Get the longest increasing and decreasing subsequences
	lis := longestIncreasingSubsequence(permutation)
	lds := longestDecreasingSubsequence(permutation)

	// Print the results
	fmt.Println(strings.Trim(fmt.Sprint(lis), "[]"))
	fmt.Println(strings.Trim(fmt.Sprint(lds), "[]"))
}

// Function to get the longest increasing subsequence
func longestIncreasingSubsequence(arr []int) []int {
	n := len(arr)
	dp := make([]int, n)
	prev := make([]int, n)
	for i := range prev {
		prev[i] = -1
	}

	lisLength := 0
	lisEnd := 0

	// Find the length of LIS and store the previous index for each element
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > lisLength {
			lisLength = dp[i]
			lisEnd = i
		}
	}

	// Reconstruct the LIS
	lis := []int{}
	for lisEnd != -1 {
		lis = append([]int{arr[lisEnd]}, lis...)
		lisEnd = prev[lisEnd]
	}

	return lis
}

// Function to get the longest decreasing subsequence
func longestDecreasingSubsequence(arr []int) []int {
	n := len(arr)
	dp := make([]int, n)
	prev := make([]int, n)
	for i := range prev {
		prev[i] = -1
	}

	ldsLength := 0
	ldsEnd := 0

	// Find the length of LDS and store the previous index for each element
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > ldsLength {
			ldsLength = dp[i]
			ldsEnd = i
		}
	}

	// Reconstruct the LDS
	lds := []int{}
	for ldsEnd != -1 {
		lds = append([]int{arr[ldsEnd]}, lds...)
		ldsEnd = prev[ldsEnd]
	}

	return lds
}
