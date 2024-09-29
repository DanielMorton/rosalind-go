package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FindSubsequenceIndices finds one collection of indices of s where t appears as a subsequence
func FindSubsequenceIndices(s, t string) []int {
	var indices []int
	tIndex := 0

	for i := 0; i < len(s); i++ {
		if tIndex < len(t) && s[i] == t[tIndex] {
			indices = append(indices, i+1) // +1 to convert 0-based to 1-based index
			tIndex++
		}
		if tIndex == len(t) {
			break
		}
	}

	return indices
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_sseq.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var s, t string
	var currentString *string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, ">") {
			// Switch between sequences when encountering a header line
			if s != "" && t != "" {
				break
			}
			if s == "" {
				currentString = &s
			} else {
				currentString = &t
			}
		} else {
			// Append the line to the current string
			*currentString += line
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Find the subsequence indices
	indices := FindSubsequenceIndices(s, t)

	// Print the result
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(indices)), " "), "[]"))
}
