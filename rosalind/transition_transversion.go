package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CountTransitionsAndTransversions counts transitions and transversions between s1 and s2
func CountTransitionsAndTransversions(s1, s2 string) (int, int) {
	var transitions, transversions int

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			switch {
			case (s1[i] == 'A' && s2[i] == 'G') || (s1[i] == 'G' && s2[i] == 'A') ||
				(s1[i] == 'C' && s2[i] == 'T') || (s1[i] == 'T' && s2[i] == 'C'):
				transitions++
			default:
				transversions++
			}
		}
	}

	return transitions, transversions
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_tran.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var s1, s2 string
	var currentString *string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, ">") {
			// Switch between sequences when encountering a header line
			if s1 != "" && s2 != "" {
				break
			}
			if s1 == "" {
				currentString = &s1
			} else {
				currentString = &s2
			}
		} else {
			// Append the line to the current string
			*currentString += line
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Count transitions and transversions
	transitions, transversions := CountTransitionsAndTransversions(s1, s2)

	// Calculate the transition/transversion ratio
	ratio := float64(transitions) / float64(transversions)

	// Print the result with 11 decimal places
	fmt.Printf("%.11f\n", ratio)
}
