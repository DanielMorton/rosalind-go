package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var codonTable = map[string]string{
	"TTT": "F", "TTC": "F", "TTA": "L", "TTG": "L",
	"CTT": "L", "CTC": "L", "CTA": "L", "CTG": "L",
	"ATT": "I", "ATC": "I", "ATA": "I", "ATG": "M",
	"GTT": "V", "GTC": "V", "GTA": "V", "GTG": "V",
	"TCT": "S", "TCC": "S", "TCA": "S", "TCG": "S",
	"CCT": "P", "CCC": "P", "CCA": "P", "CCG": "P",
	"ACT": "T", "ACC": "T", "ACA": "T", "ACG": "T",
	"GCT": "A", "GCC": "A", "GCA": "A", "GCG": "A",
	"TAT": "Y", "TAC": "Y", "TAA": "Stop", "TAG": "Stop",
	"CAT": "H", "CAC": "H", "CAA": "Q", "CAG": "Q",
	"AAT": "N", "AAC": "N", "AAA": "K", "AAG": "K",
	"GAT": "D", "GAC": "D", "GAA": "E", "GAG": "E",
	"TGT": "C", "TGC": "C", "TGA": "Stop", "TGG": "W",
	"CGT": "R", "CGC": "R", "CGA": "R", "CGG": "R",
	"AGT": "S", "AGC": "S", "AGA": "R", "AGG": "R",
	"GGT": "G", "GGC": "G", "GGA": "G", "GGG": "G",
}

// Function to get reverse complement of a DNA string
func reverseComplement(dna string) string {
	var complement = map[byte]byte{
		'A': 'T', 'T': 'A', 'C': 'G', 'G': 'C',
	}
	n := len(dna)
	reverse := make([]byte, n)
	for i := 0; i < n; i++ {
		reverse[n-i-1] = complement[dna[i]]
	}
	return string(reverse)
}

// Function to translate a DNA sequence into a protein string
func translate(dna string) string {
	var protein strings.Builder
	for i := 0; i+3 <= len(dna); i += 3 {
		codon := dna[i : i+3]
		aminoAcid := codonTable[codon]
		if aminoAcid == "Stop" {
			break
		}
		protein.WriteString(aminoAcid)
	}
	return protein.String()
}

// Function to find ORFs in a DNA sequence
func findORFs(dna string) map[string]bool {
	proteins := make(map[string]bool)
	n := len(dna)
	for frame := 0; frame < 3; frame++ {
		for i := frame; i+3 <= n; i += 3 {
			if dna[i:i+3] == "ATG" {
				protein := translate(dna[i:])
				if protein != "" {
					proteins[protein] = true
				}
			}
		}
	}
	return proteins
}

func main() {
	// Read input file
	data, err := ioutil.ReadFile("rosalind_orf.txt")
	if err != nil {
		panic(err)
	}

	// Process input file to extract the DNA sequence
	content := strings.Split(string(data), "\n")
	var dna strings.Builder
	for _, line := range content {
		if !strings.HasPrefix(line, ">") {
			dna.WriteString(strings.TrimSpace(line))
		}
	}

	dnaStr := dna.String()
	reverseDnaStr := reverseComplement(dnaStr)

	// Find ORFs in both DNA and its reverse complement
	orfs := findORFs(dnaStr)
	reverseOrfs := findORFs(reverseDnaStr)

	// Print distinct protein strings from ORFs
	for protein := range orfs {
		fmt.Println(protein)
	}
	for protein := range reverseOrfs {
		if !orfs[protein] {
			fmt.Println(protein)
		}
	}
}
