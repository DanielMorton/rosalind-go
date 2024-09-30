package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the genetic code mapping
var geneticCode = map[string]string{
	"UUU": "F", "UUC": "F", "UUA": "L", "UUG": "L",
	"UCU": "S", "UCC": "S", "UCA": "S", "UCG": "S",
	"UAU": "Y", "UAC": "Y", "UAA": "STOP", "UAG": "STOP",
	"UGU": "C", "UGC": "C", "UGA": "STOP", "UGG": "W",
	"CUU": "L", "CUC": "L", "CUA": "L", "CUG": "L",
	"CCU": "P", "CCC": "P", "CCA": "P", "CCG": "P",
	"CAU": "H", "CAC": "H", "CAA": "Q", "CAG": "Q",
	"CGU": "R", "CGC": "R", "CGA": "R", "CGG": "R",
	"AUU": "I", "AUC": "I", "AUA": "I", "AUG": "M",
	"ACU": "T", "ACC": "T", "ACA": "T", "ACG": "T",
	"AAU": "N", "AAC": "N", "AAA": "K", "AAG": "K",
	"AGU": "S", "AGC": "S", "AGA": "R", "AGG": "R",
	"GUU": "V", "GUC": "V", "GUA": "V", "GUG": "V",
	"GCU": "A", "GCC": "A", "GCA": "A", "GCG": "A",
	"GAU": "D", "GAC": "D", "GAA": "E", "GAG": "E",
	"GGU": "G", "GGC": "G", "GGA": "G", "GGG": "G",
}

func translateRNA(rna string) string {
	var peptide strings.Builder
	for i := 0; i < len(rna); i += 3 {
		if i+3 > len(rna) {
			break
		}
		codon := rna[i : i+3]
		aminoAcid, ok := geneticCode[codon]
		if !ok {
			fmt.Printf("Warning: Unknown codon %s\n", codon)
			continue
		}
		if aminoAcid == "STOP" {
			break
		}
		peptide.WriteString(aminoAcid)
	}
	return peptide.String()
}

func main() {
	// Read input from file
	file, err := os.Open("rosalind_ba4a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	rna := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Translate RNA to protein
	protein := translateRNA(rna)

	// Print the result
	fmt.Println(protein)
}
