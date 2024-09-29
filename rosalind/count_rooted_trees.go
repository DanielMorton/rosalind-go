package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const mod = 1000000

func B(n int) int {
	if n <= 1 {
		return 1
	}

	result := 1
	for i := 3; i <= 2*n-3; i += 2 {
		result = (result * i) % mod
	}

	return result
}

func main() {
	// Read input from file
	content, err := os.ReadFile("rosalind_root.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert content to integer
	nStr := strings.TrimSpace(string(content))
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return
	}

	// Calculate and print the result
	result := B(n)
	fmt.Println(result)
}
