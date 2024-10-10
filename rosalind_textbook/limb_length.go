package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func limbLength(n, j int, D [][]int) int {
	minLimbLength := 1 << 30 // A large number to start with

	for i := 0; i < n; i++ {
		if i == j {
			continue
		}
		for k := i + 1; k < n; k++ {
			if k == j {
				continue
			}
			length := (D[i][j] + D[j][k] - D[i][k]) / 2
			if length < minLimbLength {
				minLimbLength = length
			}
		}
	}

	return minLimbLength
}

func main() {
	file, err := os.Open("rosalind_ba7b.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read n
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// Read j
	scanner.Scan()
	j, _ := strconv.Atoi(scanner.Text())

	// Read distance matrix
	D := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		row := strings.Fields(scanner.Text())
		D[i] = make([]int, n)
		for k, val := range row {
			D[i][k], _ = strconv.Atoi(val)
		}
	}

	// Calculate and print limb length
	result := limbLength(n, j, D)
	fmt.Println(result)
}
