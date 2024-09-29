package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Graph map[string][]string

func main() {
	kmers := readInput("rosalind_grep.txt")
	//fmt.Printf("Read %d kmers\n", len(kmers))
	k := len(kmers[0]) - 1

	graph := buildDeBruijnGraph(kmers)
	//fmt.Printf("Graph has %d nodes\n", len(graph))

	firstKmer := kmers[0]
	cycles := findAllEulerianCycles(graph, firstKmer[:k], firstKmer[1:])
	//fmt.Printf("Found %d cycles\n", len(cycles))

	circularStrings := assembleCircularStrings(cycles, k, firstKmer)
	//fmt.Printf("Assembled %d circular strings\n", len(circularStrings))

	sort.Strings(circularStrings)
	for _, s := range circularStrings {
		fmt.Println(s)
	}
}

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var kmers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		kmer := strings.TrimSpace(scanner.Text())
		if kmer != "" {
			kmers = append(kmers, kmer)
		}
	}
	return kmers
}

func buildDeBruijnGraph(kmers []string) Graph {
	graph := make(Graph)
	for _, kmer := range kmers {
		prefix := kmer[:len(kmer)-1]
		suffix := kmer[1:]
		graph[prefix] = append(graph[prefix], suffix)
	}
	return graph
}

func findAllEulerianCycles(graph Graph, start, secondNode string) [][]string {
	var allCycles [][]string
	edgeCount := make(map[string]map[string]int)
	for node, neighbors := range graph {
		edgeCount[node] = make(map[string]int)
		for _, neighbor := range neighbors {
			edgeCount[node][neighbor]++
		}
	}

	var dfs func(node string, path []string)
	dfs = func(node string, path []string) {
		if len(path) > 1 && node == start && allEdgesUsed(edgeCount) {
			if path[1] == secondNode {
				allCycles = append(allCycles, append([]string{}, path...))
			}
			return
		}

		for neighbor, count := range edgeCount[node] {
			if count > 0 {
				edgeCount[node][neighbor]--
				dfs(neighbor, append(path, neighbor))
				edgeCount[node][neighbor]++
			}
		}
	}

	dfs(start, []string{start})

	return allCycles
}

func allEdgesUsed(edgeCount map[string]map[string]int) bool {
	for _, neighbors := range edgeCount {
		for _, count := range neighbors {
			if count > 0 {
				return false
			}
		}
	}
	return true
}

func assembleCircularStrings(cycles [][]string, k int, firstKmer string) []string {
	var circularStrings []string
	seen := make(map[string]bool)
	for _, cycle := range cycles {
		s := firstKmer
		for i := 2; i < len(cycle); i++ {
			s += cycle[i][k-1:]
		}
		// Trim the string to remove the wrapped-around part
		s = s[:len(s)-k]
		if !seen[s] {
			circularStrings = append(circularStrings, s)
			seen[s] = true
		}
	}
	return circularStrings
}
