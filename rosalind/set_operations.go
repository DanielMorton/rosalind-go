package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Function to parse a set from a string
func parseSet(setStr string) map[int]struct{} {
	setStr = strings.Trim(setStr, "{}")
	elements := strings.Split(setStr, ", ")
	set := make(map[int]struct{})
	for _, elem := range elements {
		if elem != "" {
			num, _ := strconv.Atoi(elem)
			set[num] = struct{}{}
		}
	}
	return set
}

// Function to perform set union
func setUnion(A, B map[int]struct{}) map[int]struct{} {
	unionSet := make(map[int]struct{})
	for key := range A {
		unionSet[key] = struct{}{}
	}
	for key := range B {
		unionSet[key] = struct{}{}
	}
	return unionSet
}

// Function to perform set intersection
func setIntersection(A, B map[int]struct{}) map[int]struct{} {
	intersectionSet := make(map[int]struct{})
	for key := range A {
		if _, found := B[key]; found {
			intersectionSet[key] = struct{}{}
		}
	}
	return intersectionSet
}

// Function to perform set difference A - B
func setDifference(A, B map[int]struct{}) map[int]struct{} {
	differenceSet := make(map[int]struct{})
	for key := range A {
		if _, found := B[key]; !found {
			differenceSet[key] = struct{}{}
		}
	}
	return differenceSet
}

// Function to compute set complement
func setComplement(A map[int]struct{}, n int) map[int]struct{} {
	complementSet := make(map[int]struct{})
	for i := 1; i <= n; i++ {
		if _, found := A[i]; !found {
			complementSet[i] = struct{}{}
		}
	}
	return complementSet
}

// Function to print the set
func printSet(set map[int]struct{}) {
	var elements []int
	for key := range set {
		elements = append(elements, key)
	}
	fmt.Printf("{")
	for i, elem := range elements {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%d", elem)
	}
	fmt.Printf("}\n")
}

func main() {
	// Read the input file
	data, err := ioutil.ReadFile("rosalind_seto.txt")
	if err != nil {
		panic(err)
	}

	// Split the file content into lines
	lines := strings.Split(string(data), "\n")

	// Read the universal set size
	n, _ := strconv.Atoi(lines[0])

	// Read sets A and B
	A := parseSet(lines[1])
	B := parseSet(lines[2])

	// Perform the set operations
	unionSet := setUnion(A, B)
	intersectionSet := setIntersection(A, B)
	differenceAB := setDifference(A, B)
	differenceBA := setDifference(B, A)
	complementA := setComplement(A, n)
	complementB := setComplement(B, n)

	// Print the results
	printSet(unionSet)
	printSet(intersectionSet)
	printSet(differenceAB)
	printSet(differenceBA)
	printSet(complementA)
	printSet(complementB)
}
