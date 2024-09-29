package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// choose calculates n choose k
func choose(n, k int) int {
	if k > n {
		return 0
	}
	if k > n/2 {
		k = n - k
	}
	result := 1
	for i := 1; i <= k; i++ {
		result *= n - k + i
		result /= i
	}
	return result
}

func countQuartets(n int) int {
	return choose(n, 4)
}

func main() {
	file, err := os.Open("rosalind_cntq.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	result := countQuartets(n)
	fmt.Println(result % 1000000)
}
