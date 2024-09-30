package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var geneticCode = map[string]string{
	"TTT": "F", "TTC": "F", "TTA": "L", "TTG": "L",
	"TCT": "S", "TCC": "S", "TCA": "S", "TCG": "S",
	"TAT": "Y", "TAC": "Y", "TAA": "STOP", "TAG": "STOP",
	"TGT": "C", "TGC": "C", "TGA": "STOP", "TGG": "W",
	"CTT": "L", "CTC": "L", "CTA": "L", "CTG": "L",
	"CCT": "P", "CCC": "P", "CCA": "P", "CCG": "P",
	"CAT": "H", "CAC": "H", "CAA": "Q", "CAG": "Q",
	"CGT": "R", "CGC": "R", "CGA": "R", "CGG": "R",
	"ATT": "I", "ATC": "I", "ATA": "I", "ATG": "M",
	"ACT": "T", "ACC": "T", "ACA": "T", "ACG": "T",
	"AAT": "N", "AAC": "N", "AAA": "K", "AAG": "K",
	"AGT": "S", "AGC": "S", "AGA": "R", "AGG": "R",
	"GTT": "V", "GTC": "V", "GTA": "V", "GTG": "V",
	"GCT": "A", "GCC": "A", "GCA": "A", "GCG": "A",
	"GAT": "D", "GAC": "D", "GAA": "E", "GAG": "E",
	"GGT": "G", "GGC": "G", "GGA": "G", "GGG": "G",
}

func translateDNA(dna string) string {
	var peptide strings.Builder
	for i := 0; i < len(dna); i += 3 {
		if i+3 > len(dna) {
			break
		}
		codon := dna[i : i+3]
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

func reverseComplement(dna string) string {
	complement := map[byte]byte{'A': 'T', 'T': 'A', 'C': 'G', 'G': 'C'}
	reversed := make([]byte, len(dna))
	for i := 0; i < len(dna); i++ {
		reversed[len(dna)-1-i] = complement[dna[i]]
	}
	return string(reversed)
}

func findEncodingSubstrings(dna, peptide string) []string {
	var result []string
	peptideLength := len(peptide) * 3

	for i := 0; i < 3; i++ {
		for j := i; j <= len(dna)-peptideLength; j += 3 {
			substring := dna[j : j+peptideLength]
			if translateDNA(substring) == peptide {
				result = append(result, substring)
			}
			revComp := reverseComplement(substring)
			if translateDNA(revComp) == peptide {
				result = append(result, substring)
			}
		}
	}

	return result
}

func main() {
	file, err := os.Open("rosalind_ba4b.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	dna := scanner.Text()
	scanner.Scan()
	peptide := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	encodingSubstrings := findEncodingSubstrings(dna, peptide)

	for _, substring := range encodingSubstrings {
		fmt.Println(substring)
	}
}
