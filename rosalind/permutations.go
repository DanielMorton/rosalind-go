package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadNumber reads the single integer n from the input file.
func ReadNumber(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return 0, fmt.Errorf("no data found")
	}

	num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return 0, err
	}

	return num, nil
}

// Permute generates all permutations of a slice of integers.
func Permute(arr []int, start int, result *[][]int) {
	if start == len(arr) {
		perm := make([]int, len(arr))
		copy(perm, arr)
		*result = append(*result, perm)
		return
	}
	for i := start; i < len(arr); i++ {
		arr[start], arr[i] = arr[i], arr[start]
		Permute(arr, start+1, result)
		arr[start], arr[i] = arr[i], arr[start] // backtrack
	}
}

// GeneratePermutations generates all permutations of numbers 1 to n.
func GeneratePermutations(n int) [][]int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	var result [][]int
	Permute(nums, 0, &result)
	return result
}

func main() {
	n, err := ReadNumber("rosalind_perm.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	permutations := GeneratePermutations(n)
	fmt.Println(len(permutations))
	for _, perm := range permutations {
		fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(perm)), " "), "[]"))
	}
}
