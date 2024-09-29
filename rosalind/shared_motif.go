package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadDNA reads DNA strings from the input file and returns them as a slice of strings.
func ReadDNA(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sequences []string
	var sequence string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, ">") {
			if sequence != "" {
				sequences = append(sequences, sequence)
				sequence = ""
			}
		} else {
			sequence += line
		}
	}
	if sequence != "" {
		sequences = append(sequences, sequence)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sequences, nil
}

// LongestCommonSubstring finds the longest common substring among a list of DNA sequences.
func LongestCommonSubstring(sequences []string) string {
	if len(sequences) == 0 {
		return ""
	}

	first := sequences[0]
	longest := ""

	for length := 1; length <= len(first); length++ {
		for start := 0; start <= len(first)-length; start++ {
			substring := first[start : start+length]
			found := true
			for _, seq := range sequences[1:] {
				if !strings.Contains(seq, substring) {
					found = false
					break
				}
			}
			if found && length > len(longest) {
				longest = substring
			}
		}
	}

	return longest
}

func main() {
	sequences, err := ReadDNA("rosalind_lcsm.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lcs := LongestCommonSubstring(sequences)
	fmt.Println(lcs)
}
