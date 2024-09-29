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

	sequences := make([]string, 0)
	var currentSequence strings.Builder
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, ">") {
			if currentSequence.Len() > 0 {
				sequences = append(sequences, currentSequence.String())
				currentSequence.Reset()
			}
		} else {
			currentSequence.WriteString(line)
		}
	}

	if currentSequence.Len() > 0 {
		sequences = append(sequences, currentSequence.String())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sequences, nil
}

func fittingAlignment(s, t string) (int, string, string) {
	m, n := len(s), len(t)

	// Initialize the score matrix
	score := make([][]int, m+1)
	for i := range score {
		score[i] = make([]int, n+1)
	}

	// Initialize the traceback matrix
	traceback := make([][]int, m+1)
	for i := range traceback {
		traceback[i] = make([]int, n+1)
	}

	// Fill the first row with zeros (no penalty for starting anywhere in s)
	for i := 1; i <= m; i++ {
		score[i][0] = 0
		traceback[i][0] = 1 // Up
	}

	// Fill the first column with increasing penalties
	for j := 1; j <= n; j++ {
		score[0][j] = -j
		traceback[0][j] = 2 // Left
	}

	// Fill the rest of the matrix
	maxScore, maxI, maxJ := 0, 0, 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			match := score[i-1][j-1]
			if s[i-1] == t[j-1] {
				match++
			} else {
				match--
			}
			delete := score[i-1][j] - 1
			insert := score[i][j-1] - 1

			score[i][j] = max(match, delete, insert)

			if score[i][j] == match {
				traceback[i][j] = 0 // Diagonal
			} else if score[i][j] == delete {
				traceback[i][j] = 1 // Up
			} else {
				traceback[i][j] = 2 // Left
			}

			if j == n && score[i][j] > maxScore {
				maxScore = score[i][j]
				maxI, maxJ = i, j
			}
		}
	}

	// Traceback
	var alignS, alignT strings.Builder
	i, j := maxI, maxJ
	for j > 0 {
		if traceback[i][j] == 0 { // Diagonal
			alignS.WriteByte(s[i-1])
			alignT.WriteByte(t[j-1])
			i--
			j--
		} else if traceback[i][j] == 1 { // Up
			alignS.WriteByte(s[i-1])
			alignT.WriteByte('-')
			i--
		} else { // Left
			alignS.WriteByte('-')
			alignT.WriteByte(t[j-1])
			j--
		}
	}

	return maxScore, reverse(alignS.String()), reverse(alignT.String())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}

func main() {
	filename := "rosalind_sims.txt"
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
	score, alignS, alignT := fittingAlignment(s, t)

	fmt.Println(score)
	fmt.Println(alignS)
	fmt.Println(alignT)
}
