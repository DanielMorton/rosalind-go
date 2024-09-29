package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Function to parse the input file and return the pairs of permutations
func readPermutations(filename string) [][2][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var permutations [][2][]int
	var pair [2][]int
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		if len(nums) == 0 {
			continue
		}
		perm := make([]int, len(nums))
		for i, num := range nums {
			n, _ := strconv.Atoi(num)
			perm[i] = n
		}
		pair[count] = perm
		count++
		if count == 2 {
			permutations = append(permutations, pair)
			pair = [2][]int{}
			count = 0
		}
	}

	if count != 0 {
		log.Fatal("Input file does not contain an even number of permutations")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return permutations
}

// Function to calculate the reversal distance using bidirectional BFS
func reversalDistance(start, target []int) int {
	n := len(start)
	if isEqual(start, target) {
		return 0
	}

	// Bidirectional BFS initialization
	visitedStart := map[string]bool{permToString(start): true}
	visitedTarget := map[string]bool{permToString(target): true}
	queueStart := [][]int{start}
	queueTarget := [][]int{target}
	level := 0

	for len(queueStart) > 0 && len(queueTarget) > 0 {
		// Process the start queue
		nextQueueStart := [][]int{}
		for _, perm := range queueStart {
			for i := 0; i < n-1; i++ {
				for j := i + 1; j < n; j++ {
					newPerm := reverse(perm, i, j)
					newPermStr := permToString(newPerm)
					if visitedTarget[newPermStr] {
						return level + 1
					}
					if !visitedStart[newPermStr] {
						visitedStart[newPermStr] = true
						nextQueueStart = append(nextQueueStart, newPerm)
					}
				}
			}
		}
		queueStart = nextQueueStart
		level++

		// Process the target queue
		nextQueueTarget := [][]int{}
		for _, perm := range queueTarget {
			for i := 0; i < n-1; i++ {
				for j := i + 1; j < n; j++ {
					newPerm := reverse(perm, i, j)
					newPermStr := permToString(newPerm)
					if visitedStart[newPermStr] {
						return level + 1
					}
					if !visitedTarget[newPermStr] {
						visitedTarget[newPermStr] = true
						nextQueueTarget = append(nextQueueTarget, newPerm)
					}
				}
			}
		}
		queueTarget = nextQueueTarget
		level++
	}

	return -1 // should not happen
}

// Helper function to reverse a segment of the permutation
func reverse(perm []int, i, j int) []int {
	newPerm := make([]int, len(perm))
	copy(newPerm, perm)
	for k := 0; k <= (j-i)/2; k++ {
		newPerm[i+k], newPerm[j-k] = perm[j-k], perm[i+k]
	}
	return newPerm
}

// Helper function to check if two permutations are equal
func isEqual(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Helper function to convert a permutation to a string for hashing purposes
func permToString(perm []int) string {
	str := ""
	for _, num := range perm {
		str += strconv.Itoa(num) + " "
	}
	return str
}

func main() {
	// Read the input from the file
	pairs := readPermutations("rosalind_rear.txt")

	// Compute and output the reversal distance for each pair of permutations
	for _, pair := range pairs {
		start := pair[0]
		target := pair[1]
		distance := reversalDistance(start, target)
		fmt.Printf("%d ", distance)
	}
	fmt.Println()
}
