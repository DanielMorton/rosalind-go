package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read input from file
	content, err := ioutil.ReadFile("rosalind_pdpl.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse input
	input := strings.TrimSpace(string(content))
	numbers := strings.Fields(input)
	L := make([]int, len(numbers))
	for i, num := range numbers {
		L[i], _ = strconv.Atoi(num)
	}

	fmt.Println("Input L:", L)

	// Solve the Partial Digest Problem
	X := solvePartialDigest(L)

	// Print the result
	fmt.Println("Output X:", X)
}

func solvePartialDigest(L []int) []int {
	sort.Ints(L)
	width := L[len(L)-1]
	X := []int{0, width}
	D := make([]int, len(L)-1)
	copy(D, L[:len(L)-1])
	result := place(D, X, width)
	sort.Ints(result)
	return result
}

func place(D, X []int, width int) []int {
	if len(D) == 0 {
		return X
	}

	y := D[len(D)-1]

	// Try placing y
	if isConsistent(D, X, y) {
		newX := append([]int{}, X...)
		newX = append(newX, y)
		newD := removeDistances(D, y, newX)
		result := place(newD, newX, width)
		if result != nil {
			return result
		}
	}

	// Try placing width-y
	if y != width-y && isConsistent(D, X, width-y) {
		newX := append([]int{}, X...)
		newX = append(newX, width-y)
		newD := removeDistances(D, width-y, newX)
		result := place(newD, newX, width)
		if result != nil {
			return result
		}
	}

	return nil
}

func isConsistent(D, X []int, y int) bool {
	for _, x := range X {
		diff := abs(x - y)
		if !contains(D, diff) && diff != 0 {
			return false
		}
	}
	return true
}

func removeDistances(D []int, y int, X []int) []int {
	newD := make([]int, 0)
	for _, d := range D {
		keep := true
		for _, x := range X {
			if d == abs(x-y) {
				keep = false
				break
			}
		}
		if keep {
			newD = append(newD, d)
		}
	}
	return newD
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
