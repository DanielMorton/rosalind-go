package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func greedySorting(perm []int) [][]int {
	steps := [][]int{}
	for k := 1; k <= len(perm); k++ {
		if abs(perm[k-1]) != k || perm[k-1] < 0 {
			// Find the index of k or -k
			idx := k - 1
			for i := k - 1; i < len(perm); i++ {
				if abs(perm[i]) == k {
					idx = i
					break
				}
			}
			// Reverse the segment
			reverse(perm, k-1, idx)
			steps = append(steps, append([]int{}, perm...))

			// If the number is still negative, flip its sign
			if perm[k-1] < 0 {
				perm[k-1] = -perm[k-1]
				steps = append(steps, append([]int{}, perm...))
			}
		}
	}
	return steps
}

func reverse(perm []int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		perm[i], perm[j] = -perm[j], -perm[i]
	}
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func formatPermutation(perm []int) string {
	strs := make([]string, len(perm))
	for i, num := range perm {
		if num > 0 {
			strs[i] = "+" + strconv.Itoa(num)
		} else {
			strs[i] = strconv.Itoa(num)
		}
	}
	return "(" + strings.Join(strs, " ") + ")"
}

func main() {
	file, err := os.Open("rosalind_ba6a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	// Parse the input permutation
	line = strings.Trim(line, "()")
	parts := strings.Fields(line)
	perm := make([]int, len(parts))
	for i, part := range parts {
		perm[i], _ = strconv.Atoi(strings.TrimPrefix(part, "+"))
	}

	// Apply GreedySorting
	steps := greedySorting(perm)

	// Print the results
	for _, step := range steps {
		fmt.Println(formatPermutation(step))
	}
}
