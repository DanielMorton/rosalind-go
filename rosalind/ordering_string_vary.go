package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Generate all strings of length up to maxLength from the given alphabet
func generateStrings(alphabet []string, maxLength int) []string {
	var result []string
	var backtrack func(current string, length int)

	backtrack = func(current string, length int) {
		if length > maxLength {
			return
		}
		if length > 0 {
			result = append(result, current)
		}
		for _, symbol := range alphabet {
			backtrack(current+symbol, length+1)
		}
	}

	backtrack("", 0)
	return result
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_lexv.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the alphabet
	var alphabet []string
	if scanner.Scan() {
		alphabet = strings.Fields(scanner.Text())
	}

	// Read the maximum length
	var maxLength int
	if scanner.Scan() {
		maxLength, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
	}

	// Generate all strings
	strings := generateStrings(alphabet, maxLength)

	// Print the results
	for _, str := range strings {
		fmt.Println(str)
	}
}
