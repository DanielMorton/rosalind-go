package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// AminoAcid represents an amino acid with its mass and symbol
type AminoAcid struct {
	Mass   int
	Symbol string
}

// Define the standard amino acid mass table
var aminoAcids = []AminoAcid{
	{71, "A"}, {103, "C"}, {115, "D"}, {129, "E"}, {147, "F"},
	{57, "G"}, {137, "H"}, {113, "I"}, {128, "K"}, {113, "L"},
	{131, "M"}, {114, "N"}, {97, "P"}, {128, "Q"}, {156, "R"},
	{87, "S"}, {101, "T"}, {99, "V"}, {186, "W"}, {163, "Y"},
}

func main() {
	// Read input from file
	spectrum := readSpectrum("rosalind_ba11a.txt")

	// Construct the graph
	graph := constructGraph(spectrum)

	// Print the graph
	for _, edge := range graph {
		fmt.Printf("%d->%d:%s\n", edge.From, edge.To, edge.Label)
	}
}

func readSpectrum(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	var spectrum []int
	for _, s := range strings.Fields(line) {
		mass, _ := strconv.Atoi(s)
		spectrum = append(spectrum, mass)
	}

	sort.Ints(spectrum)
	return spectrum
}

type Edge struct {
	From  int
	To    int
	Label string
}

func constructGraph(spectrum []int) []Edge {
	var graph []Edge

	// Include 0 in the spectrum if it's not already there
	if spectrum[0] != 0 {
		spectrum = append([]int{0}, spectrum...)
	}

	for i := 0; i < len(spectrum); i++ {
		for j := i + 1; j < len(spectrum); j++ {
			diff := spectrum[j] - spectrum[i]
			label := findAminoAcid(diff)
			if label != "" {
				graph = append(graph, Edge{spectrum[i], spectrum[j], label})
			}
		}
	}

	// Sort the graph edges
	sort.Slice(graph, func(i, j int) bool {
		if graph[i].From != graph[j].From {
			return graph[i].From < graph[j].From
		}
		return graph[i].To < graph[j].To
	})

	return graph
}

func findAminoAcid(mass int) string {
	for _, aa := range aminoAcids {
		if aa.Mass == mass {
			return aa.Symbol
		}
	}
	return ""
}
