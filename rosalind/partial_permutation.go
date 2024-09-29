package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PartialPermutations calculates P(n, k) = n * (n-1) * ... * (n-k+1) mod 1,000,000
func PartialPermutations(n, k int) int {
	result := 1
	mod := 1000000
	for i := 0; i < k; i++ {
		result = (result * (n - i)) % mod
	}
	return result
}

// Read integers from the input file
func readInput(filename string) (int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	var n, k int
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return 0, 0, fmt.Errorf("invalid input format")
		}
		n, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, err
		}
		k, err = strconv.Atoi(parts[1])
		if err != nil {
			return 0, 0, err
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}
	return n, k, nil
}

func main() {
	n, k, err := readInput("rosalind_pper.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	result := PartialPermutations(n, k)
	fmt.Println(result)
}
