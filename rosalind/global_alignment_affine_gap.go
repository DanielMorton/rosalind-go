package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	gapOpenPenalty   = 11
	gapExtendPenalty = 1
	negativeInfinity = -1000000
)

var blosum62 = map[string]map[string]int{
	"A": {"A": 4, "C": 0, "D": -2, "E": -1, "F": -2, "G": 0, "H": -2, "I": -1, "K": -1, "L": -1, "M": -1, "N": -2, "P": -1, "Q": -1, "R": -1, "S": 1, "T": 0, "V": 0, "W": -3, "Y": -2},
	"C": {"A": 0, "C": 9, "D": -3, "E": -4, "F": -2, "G": -3, "H": -3, "I": -1, "K": -3, "L": -1, "M": -1, "N": -3, "P": -3, "Q": -3, "R": -3, "S": -1, "T": -1, "V": -1, "W": -2, "Y": -2},
	"D": {"A": -2, "C": -3, "D": 6, "E": 2, "F": -3, "G": -1, "H": -1, "I": -3, "K": -1, "L": -4, "M": -3, "N": 1, "P": -1, "Q": 0, "R": -2, "S": 0, "T": -1, "V": -3, "W": -4, "Y": -3},
	"E": {"A": -1, "C": -4, "D": 2, "E": 5, "F": -3, "G": -2, "H": 0, "I": -3, "K": 1, "L": -3, "M": -2, "N": 0, "P": -1, "Q": 2, "R": 0, "S": 0, "T": -1, "V": -2, "W": -3, "Y": -2},
	"F": {"A": -2, "C": -2, "D": -3, "E": -3, "F": 6, "G": -3, "H": -1, "I": 0, "K": -3, "L": 0, "M": 0, "N": -3, "P": -4, "Q": -3, "R": -3, "S": -2, "T": -2, "V": -1, "W": 1, "Y": 3},
	"G": {"A": 0, "C": -3, "D": -1, "E": -2, "F": -3, "G": 6, "H": -2, "I": -4, "K": -2, "L": -4, "M": -3, "N": 0, "P": -2, "Q": -2, "R": -2, "S": 0, "T": -2, "V": -3, "W": -2, "Y": -3},
	"H": {"A": -2, "C": -3, "D": -1, "E": 0, "F": -1, "G": -2, "H": 8, "I": -3, "K": -1, "L": -3, "M": -2, "N": 1, "P": -2, "Q": 0, "R": 0, "S": -1, "T": -2, "V": -3, "W": -2, "Y": 2},
	"I": {"A": -1, "C": -1, "D": -3, "E": -3, "F": 0, "G": -4, "H": -3, "I": 4, "K": -3, "L": 2, "M": 1, "N": -3, "P": -3, "Q": -3, "R": -3, "S": -2, "T": -1, "V": 3, "W": -3, "Y": -1},
	"K": {"A": -1, "C": -3, "D": -1, "E": 1, "F": -3, "G": -2, "H": -1, "I": -3, "K": 5, "L": -2, "M": -1, "N": 0, "P": -1, "Q": 1, "R": 2, "S": 0, "T": -1, "V": -2, "W": -3, "Y": -2},
	"L": {"A": -1, "C": -1, "D": -4, "E": -3, "F": 0, "G": -4, "H": -3, "I": 2, "K": -2, "L": 4, "M": 2, "N": -3, "P": -3, "Q": -2, "R": -2, "S": -2, "T": -1, "V": 1, "W": -2, "Y": -1},
	"M": {"A": -1, "C": -1, "D": -3, "E": -2, "F": 0, "G": -3, "H": -2, "I": 1, "K": -1, "L": 2, "M": 5, "N": -2, "P": -2, "Q": 0, "R": -1, "S": -1, "T": -1, "V": 1, "W": -1, "Y": -1},
	"N": {"A": -2, "C": -3, "D": 1, "E": 0, "F": -3, "G": 0, "H": 1, "I": -3, "K": 0, "L": -3, "M": -2, "N": 6, "P": -2, "Q": 0, "R": 0, "S": 1, "T": 0, "V": -3, "W": -4, "Y": -2},
	"P": {"A": -1, "C": -3, "D": -1, "E": -1, "F": -4, "G": -2, "H": -2, "I": -3, "K": -1, "L": -3, "M": -2, "N": -2, "P": 7, "Q": -1, "R": -2, "S": -1, "T": -1, "V": -2, "W": -4, "Y": -3},
	"Q": {"A": -1, "C": -3, "D": 0, "E": 2, "F": -3, "G": -2, "H": 0, "I": -3, "K": 1, "L": -2, "M": 0, "N": 0, "P": -1, "Q": 5, "R": 1, "S": 0, "T": -1, "V": -2, "W": -2, "Y": -1},
	"R": {"A": -1, "C": -3, "D": -2, "E": 0, "F": -3, "G": -2, "H": 0, "I": -3, "K": 2, "L": -2, "M": -1, "N": 0, "P": -2, "Q": 1, "R": 5, "S": -1, "T": -1, "V": -3, "W": -3, "Y": -2},
	"S": {"A": 1, "C": -1, "D": 0, "E": 0, "F": -2, "G": 0, "H": -1, "I": -2, "K": 0, "L": -2, "M": -1, "N": 1, "P": -1, "Q": 0, "R": -1, "S": 4, "T": 1, "V": -2, "W": -3, "Y": -2},
	"T": {"A": 0, "C": -1, "D": -1, "E": -1, "F": -2, "G": -2, "H": -2, "I": -1, "K": -1, "L": -1, "M": -1, "N": 0, "P": -1, "Q": -1, "R": -1, "S": 1, "T": 5, "V": 0, "W": -2, "Y": -2},
	"V": {"A": 0, "C": -1, "D": -3, "E": -2, "F": -1, "G": -3, "H": -3, "I": 3, "K": -2, "L": 1, "M": 1, "N": -3, "P": -2, "Q": -2, "R": -3, "S": -2, "T": 0, "V": 4, "W": -3, "Y": -1},
	"W": {"A": -3, "C": -2, "D": -4, "E": -3, "F": 1, "G": -2, "H": -2, "I": -3, "K": -3, "L": -2, "M": -1, "N": -4, "P": -4, "Q": -2, "R": -3, "S": -3, "T": -2, "V": -3, "W": 11, "Y": 2},
	"Y": {"A": -2, "C": -2, "D": -3, "E": -2, "F": 3, "G": -3, "H": 2, "I": -1, "K": -2, "L": -1, "M": -1, "N": -2, "P": -3, "Q": -1, "R": -2, "S": -2, "T": -2, "V": -1, "W": 2, "Y": 7},
}

func main() {
	// Read input from file
	s, t := readFASTA("rosalind_gaff.txt")

	// Perform alignment with affine gap penalty
	score, aligned1, aligned2 := affineGapAlign(s, t)

	// Print results
	fmt.Println(score)
	fmt.Println(aligned1)
	fmt.Println(aligned2)
}

func readFASTA(filename string) (string, string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var s, t string
	currentSeq := ""

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if currentSeq != "" {
				s = currentSeq
				currentSeq = ""
			}
		} else {
			currentSeq += line
		}
	}
	t = currentSeq

	return s, t
}

func affineGapAlign(s, t string) (int, string, string) {
	m, n := len(s), len(t)
	M := make([][][]int, m+1)
	X := make([][][]int, m+1)
	Y := make([][][]int, m+1)

	for i := range M {
		M[i] = make([][]int, n+1)
		X[i] = make([][]int, n+1)
		Y[i] = make([][]int, n+1)
		for j := range M[i] {
			M[i][j] = []int{negativeInfinity, 0, 0}
			X[i][j] = []int{negativeInfinity, 0, 0}
			Y[i][j] = []int{negativeInfinity, 0, 0}
		}
	}

	M[0][0] = []int{0, 0, 0} // score, prev matrix (0:M, 1:X, 2:Y)
	X[0][0] = []int{negativeInfinity, 0, 0}
	Y[0][0] = []int{negativeInfinity, 0, 0}

	// Initialize first row and column
	for i := 1; i <= m; i++ {
		M[i][0] = []int{negativeInfinity, 0, 0}
		X[i][0] = []int{-gapOpenPenalty - (i-1)*gapExtendPenalty, 1, 0}
		Y[i][0] = []int{negativeInfinity, 0, 0}
	}
	for j := 1; j <= n; j++ {
		M[0][j] = []int{negativeInfinity, 0, 0}
		X[0][j] = []int{negativeInfinity, 0, 0}
		Y[0][j] = []int{-gapOpenPenalty - (j-1)*gapExtendPenalty, 2, 0}
	}

	// Fill the matrices
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// M matrix
			match := M[i-1][j-1][0] + blosum62[string(s[i-1])][string(t[j-1])]
			xMatch := X[i-1][j-1][0] + blosum62[string(s[i-1])][string(t[j-1])]
			yMatch := Y[i-1][j-1][0] + blosum62[string(s[i-1])][string(t[j-1])]
			maxMatch := max(match, xMatch, yMatch)
			if maxMatch == match {
				M[i][j] = []int{maxMatch, 0, 0}
			} else if maxMatch == xMatch {
				M[i][j] = []int{maxMatch, 1, 0}
			} else {
				M[i][j] = []int{maxMatch, 2, 0}
			}

			// X matrix
			gapOpen := M[i-1][j][0] - gapOpenPenalty
			gapExtend := X[i-1][j][0] - gapExtendPenalty
			if gapOpen > gapExtend {
				X[i][j] = []int{gapOpen, 0, 1}
			} else {
				X[i][j] = []int{gapExtend, 1, 1}
			}

			// Y matrix
			gapOpen = M[i][j-1][0] - gapOpenPenalty
			gapExtend = Y[i][j-1][0] - gapExtendPenalty
			if gapOpen > gapExtend {
				Y[i][j] = []int{gapOpen, 0, 2}
			} else {
				Y[i][j] = []int{gapExtend, 2, 2}
			}
		}
	}

	// Find the maximum score
	maxScore := max(M[m][n][0], X[m][n][0], Y[m][n][0])

	// Traceback
	aligned1, aligned2 := "", ""
	i, j := m, n
	currentMatrix := 0
	if maxScore == X[m][n][0] {
		currentMatrix = 1
	} else if maxScore == Y[m][n][0] {
		currentMatrix = 2
	}

	for i > 0 || j > 0 {
		var current []int
		switch currentMatrix {
		case 0:
			current = M[i][j]
		case 1:
			current = X[i][j]
		case 2:
			current = Y[i][j]
		}

		prevMatrix := current[1]
		matrixType := current[2]

		if matrixType == 0 { // Match/Mismatch
			aligned1 = string(s[i-1]) + aligned1
			aligned2 = string(t[j-1]) + aligned2
			i--
			j--
		} else if matrixType == 1 { // Gap in t
			aligned1 = string(s[i-1]) + aligned1
			aligned2 = "-" + aligned2
			i--
		} else { // Gap in s
			aligned1 = "-" + aligned1
			aligned2 = string(t[j-1]) + aligned2
			j--
		}

		currentMatrix = prevMatrix
	}

	return maxScore, aligned1, aligned2
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}
