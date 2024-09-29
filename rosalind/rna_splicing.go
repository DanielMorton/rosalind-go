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

// Function to remove introns from the DNA string
func removeIntrons(dna string, introns []string) string {
	for _, intron := range introns {
		dna = strings.ReplaceAll(dna, intron, "")
	}
	return dna
}

func main() {
	// Read input file
	data, err := ioutil.ReadFile("rosalind_splc.txt")
	if err != nil {
		panic(err)
	}

	// Parse the FASTA format file
	lines := strings.Split(string(data), "\n")
	var introns []string
	var currentSeq strings.Builder
	firstSeq := true

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, ">") {
			// Add previous sequence to introns if it's not the first sequence
			if currentSeq.Len() > 0 && !firstSeq {
				introns = append(introns, currentSeq.String())
				currentSeq.Reset()
			}
			firstSeq = false
		} else {
			// Build the DNA string or intron
			currentSeq.WriteString(line)
		}
	}

	// Add the last sequence (which is the last intron) to introns
	if currentSeq.Len() > 0 {
		introns = append(introns, currentSeq.String())
	}

	// The first sequence is the DNA sequence; the rest are introns
	dnaStr := introns[0]
	introns = introns[1:]

	// Remove introns from the DNA string
	exons := removeIntrons(dnaStr, introns)

	// Translate the exons into a protein string
	protein := translate(exons)

	// Print the protein string
	fmt.Println(protein)
}
