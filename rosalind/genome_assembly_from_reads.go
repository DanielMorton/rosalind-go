package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reads, err := readInput("rosalind_gasm.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	result := findCircularSuperstring(reads)
	fmt.Println(result)
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reads []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reads = append(reads, strings.TrimSpace(scanner.Text()))
	}

	return reads, scanner.Err()
}

func findCircularSuperstring(reads []string) string {
	// Add reverse complements
	allReads := append(reads, reverseComplements(reads)...)

	graph := buildGraph(allReads)
	cycles := findCycles(graph)

	if len(cycles) == 2 && len(cycles[0]) == len(cycles[1]) {
		return reconstructSuperstring(cycles[0])
	}

	return ""
}

func reverseComplements(reads []string) []string {
	rc := make([]string, len(reads))
	for i, read := range reads {
		rc[i] = reverseComplement(read)
	}
	return rc
}

func buildGraph(reads []string) map[string]string {
	graph := make(map[string]string)
	for _, read1 := range reads {
		maxOverlap := 0
		var bestMatch string
		for _, read2 := range reads {
			if read1 != read2 {
				overlap := findOverlap(read1, read2)
				if overlap > maxOverlap {
					maxOverlap = overlap
					bestMatch = read2
				}
			}
		}
		if maxOverlap > 0 {
			graph[read1] = bestMatch
		}
	}
	return graph
}

func findOverlap(s1, s2 string) int {
	for i := len(s1) - 1; i > 0; i-- {
		if s1[len(s1)-i:] == s2[:i] {
			return i
		}
	}
	return 0
}

func findCycles(graph map[string]string) [][]string {
	var cycles [][]string
	visited := make(map[string]bool)

	for node := range graph {
		if !visited[node] {
			cycle := findCycle(graph, node, visited)
			if len(cycle) > 0 {
				cycles = append(cycles, cycle)
			}
		}
	}

	return cycles
}

func findCycle(graph map[string]string, start string, visited map[string]bool) []string {
	var cycle []string
	current := start
	for {
		if visited[current] {
			break
		}
		visited[current] = true
		cycle = append(cycle, current)
		next, exists := graph[current]
		if !exists {
			return nil
		}
		if next == start {
			return cycle
		}
		current = next
	}
	return nil
}

func reconstructSuperstring(cycle []string) string {
	if len(cycle) == 0 {
		return ""
	}
	superstring := cycle[0]
	for i := 1; i < len(cycle); i++ {
		overlap := findOverlap(cycle[i-1], cycle[i])
		superstring += cycle[i][overlap:]
	}
	// Remove the suffix of the last vertex that overlaps with the prefix of the first vertex
	finalOverlap := findOverlap(cycle[len(cycle)-1], cycle[0])
	return superstring[:len(superstring)-finalOverlap]
}

func reverseComplement(s string) string {
	complement := map[byte]byte{'A': 'T', 'T': 'A', 'C': 'G', 'G': 'C'}
	result := make([]byte, len(s))
	for i := range s {
		result[len(s)-1-i] = complement[s[i]]
	}
	return string(result)
}
