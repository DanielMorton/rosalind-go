package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type AminoAcid struct {
	Mass   int
	Symbol string
}

var aminoAcids = []AminoAcid{
	{71, "A"}, {103, "C"}, {115, "D"}, {129, "E"}, {147, "F"},
	{57, "G"}, {137, "H"}, {113, "I"}, {128, "K"}, {113, "L"},
	{131, "M"}, {114, "N"}, {97, "P"}, {128, "Q"}, {156, "R"},
	{87, "S"}, {101, "T"}, {99, "V"}, {186, "W"}, {163, "Y"},
}

type Edge struct {
	To    int
	Label string
}

type Graph map[int][]Edge

func main() {
	spectrum := readSpectrum("rosalind_ba11b.txt")
	peptide := decodeIdealSpectrum(spectrum)
	fmt.Println(peptide)
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

func constructGraph(spectrum []int) Graph {
	graph := make(Graph)

	// Ensure 0 is in the spectrum
	if spectrum[0] != 0 {
		spectrum = append([]int{0}, spectrum...)
	}

	for i := 0; i < len(spectrum); i++ {
		for j := i + 1; j < len(spectrum); j++ {
			diff := spectrum[j] - spectrum[i]
			label := findAminoAcid(diff)
			if label != "" {
				graph[spectrum[i]] = append(graph[spectrum[i]], Edge{spectrum[j], label})
			}
		}
	}

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

func decodeIdealSpectrum(spectrum []int) string {
	graph := constructGraph(spectrum)
	return findPeptide(graph, 0, spectrum[len(spectrum)-1], "", spectrum)
}

func findPeptide(graph Graph, current, sink int, peptide string, spectrum []int) string {
	if current == sink {
		if compareSpectrums(idealSpectrum(peptide), spectrum) {
			return peptide
		}
		return ""
	}

	for _, edge := range graph[current] {
		for _, symbol := range strings.Split(edge.Label, "/") {
			newPeptide := peptide + symbol
			result := findPeptide(graph, edge.To, sink, newPeptide, spectrum)
			if result != "" {
				return result
			}
		}
	}

	return ""
}

func idealSpectrum(peptide string) []int {
	var spectrum []int

	// Calculate prefix masses (excluding 0)
	prefixMass := 0
	for i := 0; i < len(peptide); i++ {
		prefixMass += findMass(string(peptide[i]))
		spectrum = append(spectrum, prefixMass)
	}

	// Calculate suffix masses (excluding full peptide mass)
	suffixMass := 0
	for i := len(peptide) - 1; i > 0; i-- {
		suffixMass += findMass(string(peptide[i]))
		spectrum = append(spectrum, suffixMass)
	}

	sort.Ints(spectrum)
	return spectrum
}

func findMass(symbol string) int {
	for _, aa := range aminoAcids {
		if strings.Contains(aa.Symbol, symbol) {
			return aa.Mass
		}
	}
	return 0
}

func compareSpectrums(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
