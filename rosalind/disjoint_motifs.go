package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func canInterweave(s, t, u string) bool {
	for i := 0; i <= len(s)-len(t)-len(u); i++ {
		if checkInterweave(s[i:], t, u, 0, 0, 0) {
			return true
		}
	}
	return false
}

func checkInterweave(s, t, u string, i, j, k int) bool {
	if j == len(t) && k == len(u) {
		return true
	}
	if i == len(s) {
		return false
	}

	if j < len(t) && s[i] == t[j] && checkInterweave(s, t, u, i+1, j+1, k) {
		return true
	}
	if k < len(u) && s[i] == u[k] && checkInterweave(s, t, u, i+1, j, k+1) {
		return true
	}

	return false
}

func main() {
	file, err := os.Open("rosalind_itwv.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var s string
	var patterns []string

	if scanner.Scan() {
		s = scanner.Text()
	}

	for scanner.Scan() {
		pattern := scanner.Text()
		if pattern != "" {
			patterns = append(patterns, pattern)
		}
	}

	n := len(patterns)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if canInterweave(s, patterns[i], patterns[j]) {
				matrix[i][j] = 1
			}
		}
	}

	for _, row := range matrix {
		fmt.Println(strings.Trim(fmt.Sprint(row), "[]"))
	}
}
