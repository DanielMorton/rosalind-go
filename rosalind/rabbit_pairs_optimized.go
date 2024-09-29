package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to compute the number of rabbit pairs after n months with litter size k
func rabbitPairs(n, k int) int {
	// Base cases for month 1 and month 2
	if n == 1 || n == 2 {
		return 1
	}

	// Variables to store the last two terms in the sequence
	prev1, prev2 := 1, 1

	// Calculate the number of rabbit pairs for month 3 through n
	for i := 3; i <= n; i++ {
		current := prev1 + k*prev2
		prev2 = prev1
		prev1 = current
	}

	// Return the number of rabbit pairs at month n
	return prev1
}

func main() {
	// Open the file "rosalind_fib.txt"
	file, err := os.Open("rosalind_fib.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	scanner.Scan() // Read the first line

	// Split the line into separate components
	line := scanner.Text()
	values := strings.Split(line, " ")

	// Convert strings to integers
	n, _ := strconv.Atoi(values[0]) // number of months
	k, _ := strconv.Atoi(values[1]) // litter size

	// Calculate the total number of rabbit pairs after n months
	result := rabbitPairs(n, k)

	// Print the result
	fmt.Println(result)
}
