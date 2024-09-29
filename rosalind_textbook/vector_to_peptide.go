package main

import (
	"bufio"
	"fmt"
	"os"
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
	{4, "X"}, {5, "Z"}, // Imaginary amino acids X and Z
}

func main() {
	vector := readVector("rosalind_ba11d.txt")
	peptide := vectorToPeptide(vector)
	fmt.Printf(peptide)
}

func readVector(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := strings.TrimSpace(scanner.Text())

	var vector []int
	for _, s := range strings.Fields(line) {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Invalid input: %s", s))
		}
		vector = append(vector, num)
	}
	return vector
}

func vectorToPeptide(vector []int) string {
	var peptide strings.Builder
	currentMass := 0
	for i, bit := range vector {
		if bit == 1 {
			mass := i + 1 - currentMass
			symbol := findAminoAcid(mass)
			peptide.WriteString(symbol)
			currentMass = i + 1
		}
	}
	return peptide.String()
}

func findAminoAcid(mass int) string {
	for _, aa := range aminoAcids {
		if aa.Mass == mass {
			return aa.Symbol
		}
	}
	return fmt.Sprintf("?(%d)", mass) // Return a placeholder for unknown masses
}
