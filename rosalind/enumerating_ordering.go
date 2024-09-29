package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// Generate all permutations of the numbers 1 to n
func permutations(n int) [][]int {
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}

	var permute func([]int, int)
	result := [][]int{}

	permute = func(nums []int, k int) {
		if k == len(nums) {
			perm := make([]int, len(nums))
			copy(perm, nums)
			result = append(result, perm)
			return
		}
		for i := k; i < len(nums); i++ {
			nums[k], nums[i] = nums[i], nums[k]
			permute(nums, k+1)
			nums[k], nums[i] = nums[i], nums[k]
		}
	}

	permute(nums, 0)
	return result
}

// Generate all signed permutations from a list of permutations
func signedPermutations(perms [][]int) [][]int {
	var signedPerms [][]int
	n := len(perms[0])
	signs := make([]int, n)

	// Generate all sign combinations
	signCombinations := int(math.Pow(2, float64(n)))
	for i := 0; i < signCombinations; i++ {
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 0 {
				signs[j] = -1
			} else {
				signs[j] = 1
			}
		}

		for _, perm := range perms {
			signedPerm := make([]int, n)
			for k, v := range perm {
				signedPerm[k] = v * signs[k]
			}
			signedPerms = append(signedPerms, signedPerm)
		}
	}

	return signedPerms
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_sign.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the integer n from the file
	scanner := bufio.NewScanner(file)
	var n int
	if scanner.Scan() {
		n, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
	}

	// Generate all permutations
	perms := permutations(n)
	// Generate all signed permutations
	signedPerms := signedPermutations(perms)

	// Print the total number of signed permutations
	fmt.Println(len(signedPerms))
	// Print each signed permutation
	for _, sp := range signedPerms {
		for _, num := range sp {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
