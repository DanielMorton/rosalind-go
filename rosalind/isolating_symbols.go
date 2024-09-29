package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseFasta(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sequences []string
	var currentSeq strings.Builder
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
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

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sequences, nil
}

func calculateAlignment(s, t string) (int, int) {
	m, n := len(s), len(t)
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// Calculate the matrix
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			score := -1
			if s[i] == t[j] {
				score = 1
			}

			if i > 0 && j > 0 {
				matrix[i][j] = max(matrix[i-1][j-1]+score, matrix[i-1][j]-1, matrix[i][j-1]-1)
			} else if i > 0 {
				matrix[i][j] = matrix[i-1][j] - 1
			} else if j > 0 {
				matrix[i][j] = matrix[i][j-1] - 1
			} else {
				matrix[i][j] = score
			}
		}
	}

	// Calculate the global alignment score
	globalScore := matrix[m-1][n-1]
	for i := 0; i < m-1; i++ {
		globalScore = max(globalScore, matrix[i][n-1]-(m-1-i))
	}
	for j := 0; j < n-1; j++ {
		globalScore = max(globalScore, matrix[m-1][j]-(n-1-j))
	}

	// Calculate the sum of all matrix elements
	sum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum += matrix[i][j]
		}
	}

	return globalScore, sum
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	maxNum := nums[0]
	for _, num := range nums[1:] {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func main() {
	filename := "rosalind_osym.txt"
	sequences, err := parseFasta(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(sequences) < 2 {
		fmt.Println("Not enough sequences in the input file")
		return
	}

	s, t := sequences[0], sequences[1]
	globalScore, matrixSum := calculateAlignment(s, t)

	fmt.Println(globalScore)
	fmt.Println(matrixSum)
}
