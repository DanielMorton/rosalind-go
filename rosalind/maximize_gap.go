package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFASTA(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sequences []string
	var currentSeq strings.Builder
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if currentSeq.Len() > 0 {
				sequences = append(sequences, currentSeq.String())
				currentSeq.Reset()
			}
		} else {
			currentSeq.WriteString(line)
		}
	}

	if currentSeq.Len() > 0 {
		sequences = append(sequences, currentSeq.String())
	}

	return sequences, scanner.Err()
}

func maxGaps(s, t string) int {
	return abs(len(s) - len(t))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	sequences, err := readFASTA("rosalind_mgap.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(sequences) != 2 {
		fmt.Println("Error: Expected 2 sequences in the input file")
		return
	}

	s, t := sequences[0], sequences[1]
	result := maxGaps(s, t)
	fmt.Println(result)
}
