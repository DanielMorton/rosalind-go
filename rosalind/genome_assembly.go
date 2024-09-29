package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findLargestOverlap(s1, s2 string) (int, bool) {
	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}
	for i := minLen; i > 0; i-- {
		if s1[len(s1)-i:] == s2[:i] {
			return i, true
		}
		if s2[len(s2)-i:] == s1[:i] {
			return i, false
		}
	}
	return 0, false
}

func mergeDNAStrings(dnaStrings []string) string {
	for len(dnaStrings) > 1 {
		bestOverlap := 0
		bestI, bestJ := 0, 0
		s1First := false

		for i := 0; i < len(dnaStrings); i++ {
			for j := i + 1; j < len(dnaStrings); j++ {
				overlap, isS1First := findLargestOverlap(dnaStrings[i], dnaStrings[j])
				if overlap > bestOverlap {
					bestOverlap = overlap
					bestI, bestJ = i, j
					s1First = isS1First
				}
			}
		}

		var merged string
		if s1First {
			merged = dnaStrings[bestI] + dnaStrings[bestJ][bestOverlap:]
		} else {
			merged = dnaStrings[bestJ] + dnaStrings[bestI][bestOverlap:]
		}

		// Replace one string with merged result and remove the other
		dnaStrings[bestI] = merged
		dnaStrings = append(dnaStrings[:bestJ], dnaStrings[bestJ+1:]...)
	}

	return dnaStrings[0]
}

func main() {
	file, err := os.Open("rosalind_long.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var dnaStrings []string
	var currentDNA strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if currentDNA.Len() > 0 {
				dnaStrings = append(dnaStrings, currentDNA.String())
				currentDNA.Reset()
			}
		} else {
			currentDNA.WriteString(line)
		}
	}

	if currentDNA.Len() > 0 {
		dnaStrings = append(dnaStrings, currentDNA.String())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	superstring := mergeDNAStrings(dnaStrings)
	fmt.Println(superstring)
}
