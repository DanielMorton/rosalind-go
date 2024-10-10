package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	GAP_PENALTY = -5
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

func linearSpaceAlignment(v, w string) (int, string, string) {
	m, n := len(v), len(w)
	if m == 0 {
		return n * GAP_PENALTY, strings.Repeat("-", n), w
	}
	if n == 0 {
		return m * GAP_PENALTY, v, strings.Repeat("-", m)
	}

	alignV, alignW := "", ""
	score := 0

	var align func(int, int, int, int)
	align = func(top, bottom, left, right int) {
		if left == right {
			for i := top; i < bottom; i++ {
				alignV += string(v[i])
				alignW += "-"
				score += GAP_PENALTY
			}
			return
		}
		if top == bottom {
			for j := left; j < right; j++ {
				alignV += "-"
				alignW += string(w[j])
				score += GAP_PENALTY
			}
			return
		}

		middle := (left + right) / 2
		midNode := middleNode(v[top:bottom], w[left:right]) + top
		midEdge := middleEdge(v[top:bottom], w[left:right])

		align(top, midNode, left, middle)

		if midEdge == "↘" {
			alignV += string(v[midNode])
			alignW += string(w[middle])
			score += BLOSUM62[string(v[midNode])][string(w[middle])]
			midNode++
			middle++
		} else if midEdge == "↓" {
			alignV += string(v[midNode])
			alignW += "-"
			score += GAP_PENALTY
			midNode++
		} else { // "→"
			alignV += "-"
			alignW += string(w[middle])
			score += GAP_PENALTY
			middle++
		}

		align(midNode, bottom, middle, right)
	}

	align(0, m, 0, n)
	return score, alignV, alignW
}

func middleNode(v, w string) int {
	m, n := len(v), len(w)
	if m == 0 || n == 0 {
		return 0
	}
	mid := n / 2

	prevColumn := make([]int, m+1)
	for i := 1; i <= m; i++ {
		prevColumn[i] = prevColumn[i-1] + GAP_PENALTY
	}

	for j := 1; j <= mid; j++ {
		currentColumn := make([]int, m+1)
		currentColumn[0] = j * GAP_PENALTY
		for i := 1; i <= m; i++ {
			match := prevColumn[i-1] + BLOSUM62[string(v[i-1])][string(w[j-1])]
			delete := prevColumn[i] + GAP_PENALTY
			insert := currentColumn[i-1] + GAP_PENALTY
			currentColumn[i] = max(match, delete, insert)
		}
		prevColumn = currentColumn
	}

	fromSource := prevColumn

	prevColumn = make([]int, m+1)
	for i := 1; i <= m; i++ {
		prevColumn[i] = prevColumn[i-1] + GAP_PENALTY
	}

	for j := n; j > mid; j-- {
		currentColumn := make([]int, m+1)
		currentColumn[0] = (n - j + 1) * GAP_PENALTY
		for i := 1; i <= m; i++ {
			match := prevColumn[i-1] + BLOSUM62[string(v[i-1])][string(w[j-1])]
			delete := prevColumn[i] + GAP_PENALTY
			insert := currentColumn[i-1] + GAP_PENALTY
			currentColumn[i] = max(match, delete, insert)
		}
		prevColumn = currentColumn
	}

	fromSink := prevColumn

	maxScore, maxIndex := -1<<31, 0
	for i := 0; i <= m; i++ {
		score := fromSource[i] + fromSink[m-i]
		if score > maxScore {
			maxScore = score
			maxIndex = i
		}
	}

	return maxIndex
}

func middleEdge(v, w string) string {
	m, n := len(v), len(w)
	if m == 0 {
		return "→"
	}
	if n == 0 {
		return "↓"
	}
	mid := n / 2

	prevColumn := make([]int, m+1)
	for i := 1; i <= m; i++ {
		prevColumn[i] = prevColumn[i-1] + GAP_PENALTY
	}

	var midColumn []int
	for j := 1; j <= mid+1; j++ {
		currentColumn := make([]int, m+1)
		currentColumn[0] = j * GAP_PENALTY
		for i := 1; i <= m; i++ {
			match := prevColumn[i-1] + BLOSUM62[string(v[i-1])][string(w[j-1])]
			delete := prevColumn[i] + GAP_PENALTY
			insert := currentColumn[i-1] + GAP_PENALTY
			currentColumn[i] = max(match, delete, insert)
		}
		if j == mid {
			midColumn = make([]int, m+1)
			copy(midColumn, currentColumn)
		}
		prevColumn = currentColumn
	}

	for i := 0; i <= m; i++ {
		if i == m {
			return "→"
		}
		if midColumn[i]+GAP_PENALTY == prevColumn[i] {
			return "→"
		}
		if i < m && midColumn[i]+BLOSUM62[string(v[i])][string(w[mid])] == prevColumn[i+1] {
			return "↘"
		}
		return "↓"
	}

	return "↓" // Default case
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
	file, err := os.Open("rosalind_ba5l.txt")
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

	score, alignV, alignW := linearSpaceAlignment(v, w)
	fmt.Println(score)
	fmt.Println(alignV)
	fmt.Println(alignW)
}
