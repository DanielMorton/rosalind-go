package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var aminoAcidMass = []int{
	57, 71, 87, 97, 99, 101, 103, 113, 114, 115, 128, 129, 131, 137, 147, 156, 163, 186,
}

type Peptide struct {
	sequence []int
	mass     int
}

func expand(peptides []Peptide) []Peptide {
	var expanded []Peptide
	for _, peptide := range peptides {
		for _, mass := range aminoAcidMass {
			newSequence := make([]int, len(peptide.sequence))
			copy(newSequence, peptide.sequence)
			newSequence = append(newSequence, mass)
			newPeptide := Peptide{
				sequence: newSequence,
				mass:     peptide.mass + mass,
			}
			expanded = append(expanded, newPeptide)
		}
	}
	return expanded
}

func cyclospectrum(peptide Peptide) []int {
	if len(peptide.sequence) == 0 {
		return []int{0}
	}
	n := len(peptide.sequence)
	extendedPeptide := append(peptide.sequence, peptide.sequence...)
	spectrum := []int{0, peptide.mass} // Include 0 and total mass

	for length := 1; length < n; length++ {
		for i := 0; i < n; i++ {
			subPeptide := extendedPeptide[i : i+length]
			spectrum = append(spectrum, sum(subPeptide))
		}
	}

	sort.Ints(spectrum)
	return spectrum
}

func sum(slice []int) int {
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}

func linearSpectrum(peptide Peptide) []int {
	spectrum := []int{0}
	prefixMass := []int{0}

	for _, mass := range peptide.sequence {
		prefixMass = append(prefixMass, prefixMass[len(prefixMass)-1]+mass)
	}

	for i := 0; i < len(prefixMass); i++ {
		for j := i + 1; j < len(prefixMass); j++ {
			spectrum = append(spectrum, prefixMass[j]-prefixMass[i])
		}
	}

	sort.Ints(spectrum)
	return spectrum
}

func score(peptide Peptide, spectrum []int) int {
	theoreticalSpectrum := cyclospectrum(peptide)
	return countMatches(theoreticalSpectrum, spectrum)
}

func countMatches(spectrum1, spectrum2 []int) int {
	i, j, count := 0, 0, 0
	for i < len(spectrum1) && j < len(spectrum2) {
		if spectrum1[i] == spectrum2[j] {
			count++
			i++
			j++
		} else if spectrum1[i] < spectrum2[j] {
			i++
		} else {
			j++
		}
	}
	return count
}

func trim(leaderboard []Peptide, spectrum []int, n int) []Peptide {
	if len(leaderboard) <= n {
		return leaderboard
	}

	scoredPeptides := make([]struct {
		peptide Peptide
		score   int
	}, len(leaderboard))

	for i, peptide := range leaderboard {
		scoredPeptides[i] = struct {
			peptide Peptide
			score   int
		}{peptide, score(peptide, spectrum)}
	}

	sort.Slice(scoredPeptides, func(i, j int) bool {
		return scoredPeptides[i].score > scoredPeptides[j].score
	})

	var trimmed []Peptide
	for i := 0; i < n && i < len(scoredPeptides); i++ {
		trimmed = append(trimmed, scoredPeptides[i].peptide)
	}

	lastScore := scoredPeptides[len(trimmed)-1].score
	for i := len(trimmed); i < len(scoredPeptides) && scoredPeptides[i].score == lastScore; i++ {
		trimmed = append(trimmed, scoredPeptides[i].peptide)
	}

	return trimmed
}

func leaderboardCyclopeptideSequencing(spectrum []int, n int) Peptide {
	leaderboard := []Peptide{{sequence: []int{}, mass: 0}}
	leaderPeptide := Peptide{}
	parentMass := spectrum[len(spectrum)-1]

	for len(leaderboard) > 0 {
		leaderboard = expand(leaderboard)
		var newLeaderboard []Peptide
		for _, peptide := range leaderboard {
			if peptide.mass == parentMass {
				if score(peptide, spectrum) > score(leaderPeptide, spectrum) {
					leaderPeptide = peptide
				}
			} else if peptide.mass < parentMass {
				newLeaderboard = append(newLeaderboard, peptide)
			}
		}
		leaderboard = trim(newLeaderboard, spectrum, n)
	}

	return leaderPeptide
}

func main() {
	file, err := os.Open("rosalind_ba4g.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read N
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// Read spectrum
	scanner.Scan()
	spectrumStr := strings.Fields(scanner.Text())
	spectrum := make([]int, len(spectrumStr))
	for i, s := range spectrumStr {
		spectrum[i], _ = strconv.Atoi(s)
	}

	leaderPeptide := leaderboardCyclopeptideSequencing(spectrum, n)

	// Print the result
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(leaderPeptide.sequence)), "-"), "[]"))
}
