package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	GAP_OPENING_PENALTY   = 11
	GAP_EXTENSION_PENALTY = 1
)

var BLOSUM62 = map[string]map[string]int{
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

func affineGapAlignment(v, w string) (int, string, string) {
	m, n := len(v), len(w)

	// Create 3D slices to store the alignment scores
	M := make([][][]int, m+1)
	X := make([][][]int, m+1)
	Y := make([][][]int, m+1)
	for i := range M {
		M[i] = make([][]int, n+1)
		X[i] = make([][]int, n+1)
		Y[i] = make([][]int, n+1)
		for j := range M[i] {
			M[i][j] = make([]int, 3)
			X[i][j] = make([]int, 3)
			Y[i][j] = make([]int, 3)
		}
	}

	// Initialize the matrices
	for i := 1; i <= m; i++ {
		M[i][0][0] = -1e9 // Set to a very low value
		X[i][0][0] = -GAP_OPENING_PENALTY - (i-1)*GAP_EXTENSION_PENALTY
		Y[i][0][0] = -1e9 // Set to a very low value
	}
	for j := 1; j <= n; j++ {
		M[0][j][0] = -1e9 // Set to a very low value
		X[0][j][0] = -1e9 // Set to a very low value
		Y[0][j][0] = -GAP_OPENING_PENALTY - (j-1)*GAP_EXTENSION_PENALTY
	}

	// Fill the matrices
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			M[i][j][0] = max(M[i-1][j-1][0], X[i-1][j-1][0], Y[i-1][j-1][0]) + BLOSUM62[string(v[i-1])][string(w[j-1])]
			X[i][j][0] = max(M[i-1][j][0]-GAP_OPENING_PENALTY, X[i-1][j][0]-GAP_EXTENSION_PENALTY, Y[i-1][j][0]-GAP_OPENING_PENALTY)
			Y[i][j][0] = max(M[i][j-1][0]-GAP_OPENING_PENALTY, X[i][j-1][0]-GAP_OPENING_PENALTY, Y[i][j-1][0]-GAP_EXTENSION_PENALTY)

			// Store backtrack information
			M[i][j][1] = argmax(M[i-1][j-1][0], X[i-1][j-1][0], Y[i-1][j-1][0])
			X[i][j][1] = argmax(M[i-1][j][0]-GAP_OPENING_PENALTY, X[i-1][j][0]-GAP_EXTENSION_PENALTY, Y[i-1][j][0]-GAP_OPENING_PENALTY)
			Y[i][j][1] = argmax(M[i][j-1][0]-GAP_OPENING_PENALTY, X[i][j-1][0]-GAP_OPENING_PENALTY, Y[i][j-1][0]-GAP_EXTENSION_PENALTY)

			M[i][j][2] = 0 // M
			X[i][j][2] = 1 // X
			Y[i][j][2] = 2 // Y
		}
	}

	// Traceback
	alignV, alignW := "", ""
	i, j := m, n
	current := argmax(M[m][n][0], X[m][n][0], Y[m][n][0])
	for i > 0 || j > 0 {
		if current == 0 { // M
			alignV = string(v[i-1]) + alignV
			alignW = string(w[j-1]) + alignW
			current = M[i][j][1]
			i--
			j--
		} else if current == 1 { // X
			alignV = string(v[i-1]) + alignV
			alignW = "-" + alignW
			current = X[i][j][1]
			i--
		} else { // Y
			alignV = "-" + alignV
			alignW = string(w[j-1]) + alignW
			current = Y[i][j][1]
			j--
		}
	}

	return max(M[m][n][0], X[m][n][0], Y[m][n][0]), alignV, alignW
}

func max(a ...int) int {
	max := a[0]
	for _, v := range a[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func argmax(a ...int) int {
	maxIndex := 0
	maxValue := a[0]
	for i, v := range a[1:] {
		if v > maxValue {
			maxValue = v
			maxIndex = i + 1
		}
	}
	return maxIndex
}

func main() {
	// Read input from file
	file, err := os.Open("rosalind_ba5j.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var v, w string
	if scanner.Scan() {
		v = strings.TrimSpace(scanner.Text())
	}
	if scanner.Scan() {
		w = strings.TrimSpace(scanner.Text())
	}

	// Calculate and print the alignment
	score, alignV, alignW := affineGapAlignment(v, w)
	fmt.Println(score)
	fmt.Println(alignV)
	fmt.Println(alignW)
}
