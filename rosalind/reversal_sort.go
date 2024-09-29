package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to reverse a slice between two indices
func reverse(perm []int, i, j int) []int {
	result := make([]int, len(perm))
	copy(result, perm)
	for k := 0; k <= (j-i)/2; k++ {
		result[i+k], result[j-k] = perm[j-k], perm[i+k]
	}
	return result
}

// BFS function to find the reversal distance and the sequence of reversals
func reversalDistance(start, goal []int) (int, [][2]int) {
	queue := [][]int{start}
	visited := make(map[string]bool)
	parents := make(map[string][]int)
	moves := make(map[string][2]int)
	visited[fmt.Sprint(start)] = true
	parents[fmt.Sprint(start)] = nil

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if fmt.Sprint(current) == fmt.Sprint(goal) {
			// Backtrack to find the sequence of reversals
			var result [][2]int
			for p := current; parents[fmt.Sprint(p)] != nil; p = parents[fmt.Sprint(p)] {
				reverseStep := moves[fmt.Sprint(p)]
				result = append(result, reverseStep)
			}
			// Reverse the result since we are backtracking
			for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
				result[i], result[j] = result[j], result[i]
			}
			return len(result), result
		}

		// Explore all possible reversals
		for i := 0; i < len(current)-1; i++ {
			for j := i + 1; j < len(current); j++ {
				next := reverse(current, i, j)
				if !visited[fmt.Sprint(next)] {
					visited[fmt.Sprint(next)] = true
					queue = append(queue, next)
					parents[fmt.Sprint(next)] = current
					moves[fmt.Sprint(next)] = [2]int{i + 1, j + 1} // store 1-based indices
				}
			}
		}
	}
	return -1, nil
}

// Read input from the file and parse the permutations
func readPermutations(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var perms [][]int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		perm := make([]int, len(parts))
		for i, p := range parts {
			perm[i], _ = strconv.Atoi(p)
		}
		perms = append(perms, perm)
	}

	return perms[0], perms[1]
}

func main() {
	// Read the permutations from the input file
	start, goal := readPermutations("rosalind_sort.txt")

	// Calculate the reversal distance and the sequence of reversals
	distance, reversals := reversalDistance(start, goal)

	// Output the result
	fmt.Println(distance)
	for _, r := range reversals {
		fmt.Println(r[0], r[1])
	}
}
