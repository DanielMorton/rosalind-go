package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const epsilon = 1e-9

type Multiset map[float64]int

func parseMultiset(line string) Multiset {
	ms := make(Multiset)
	for _, s := range strings.Fields(line) {
		f, _ := strconv.ParseFloat(s, 64)
		ms[round(f)]++
	}
	return ms
}

func round(x float64) float64 {
	return math.Round(x/epsilon) * epsilon
}

func minkowskiDifference(s1, s2 Multiset) Multiset {
	result := make(Multiset)
	for x1 := range s1 {
		for x2 := range s2 {
			diff := round(x1 - x2)
			result[diff] += s1[x1] * s2[x2]
		}
	}
	return result
}

func findMaxMultiplicity(ms Multiset) (int, float64) {
	maxMult := 0
	var maxX float64
	for x, mult := range ms {
		if mult > maxMult || (mult == maxMult && math.Abs(x) < math.Abs(maxX)) {
			maxMult = mult
			maxX = x
		}
	}
	return maxMult, maxX
}

func main() {
	file, err := os.Open("rosalind_conv.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read first line
	scanner.Scan()
	s1 := parseMultiset(scanner.Text())

	// Read second line
	scanner.Scan()
	s2 := parseMultiset(scanner.Text())

	diff := minkowskiDifference(s1, s2)
	maxMult, maxX := findMaxMultiplicity(diff)

	fmt.Println(maxMult)
	fmt.Printf("%.5f\n", math.Abs(maxX))
}
