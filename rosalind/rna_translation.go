package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Complete RNA codon table mapping with all 64 codons
var codonTable = map[string]string{
	"UUU": "F", "UUC": "F", "UUA": "L", "UUG": "L",
	"UCU": "S", "UCC": "S", "UCA": "S", "UCG": "S",
	"UAU": "Y", "UAC": "Y", "UAA": "Stop", "UAG": "Stop",
	"UGU": "C", "UGC": "C", "UGA": "Stop", "UGG": "W",
	"CUU": "L", "CUC": "L", "CUA": "L", "CUG": "L",
	"CCU": "P", "CCC": "P", "CCA": "P", "CCG": "P",
	"CAU": "H", "CAC": "H", "CAA": "Q", "CAG": "Q",
	"CGU": "R", "CGC": "R", "CGA": "R", "CGG": "R",
	"AUU": "I", "AUC": "I", "AUA": "I", "AUG": "M",
	"ACU": "T", "ACC": "T", "ACA": "T", "ACG": "T",
	"AAU": "N", "AAC": "N", "AAA": "K", "AAG": "K",
	"AGA": "R", "AGG": "R",
	"AGU": "S", "AGC": "S",
	"GUU": "V", "GUC": "V", "GUA": "V", "GUG": "V",
	"GCU": "A", "GCC": "A", "GCA": "A", "GCG": "A",
	"GAU": "D", "GAC": "D", "GAA": "E", "GAG": "E",
	"GGU": "G", "GGC": "G", "GGA": "G", "GGG": "G",
}

// Function to translate RNA string to amino acid sequence
func translateRNA(rna string) string {
	var aminoAcids []string
	for i := 0; i < len(rna); i += 3 {
		if i+3 > len(rna) {
			break
		}
		codon := rna[i : i+3]
		if aminoAcid, ok := codonTable[codon]; ok {
			if aminoAcid == "Stop" {
				break
			}
			aminoAcids = append(aminoAcids, aminoAcid)
		}
	}
	return strings.Join(aminoAcids, "")
}

func main() {
	// Open the file "rosalind_prot.txt"
	file, err := os.Open("rosalind_prot.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	scanner.Scan() // Read the first line

	// The RNA string is in the first line
	rna := scanner.Text()

	// Translate the RNA string
	aminoAcidSequence := translateRNA(rna)

	// Print the result
	fmt.Println(aminoAcidSequence)
}
