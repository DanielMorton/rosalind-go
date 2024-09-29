package main

import (
	"bufio"
	"fmt"
	"os"
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
	{4, "X"}, {5, "Z"}, // Imaginary amino acids X and Z
}

func main() {
	peptide := readPeptide("rosalind_ba11c.txt")
	vector := peptideToBinaryVector(peptide)
	fmt.Println(strings.Trim(fmt.Sprint(vector), "[]"))
}

func readPeptide(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func peptideToBinaryVector(peptide string) []int {
	totalMass := 0
	for _, aa := range peptide {
		totalMass += findMass(string(aa))
	}

	vector := make([]int, totalMass)
	prefixMass := 0
	for _, aa := range peptide {
		prefixMass += findMass(string(aa))
		vector[prefixMass-1] = 1 // -1 because array is 0-indexed
	}

	return vector
}

func findMass(symbol string) int {
	for _, aa := range aminoAcids {
		if aa.Symbol == symbol {
			return aa.Mass
		}
	}
	return 0
}
