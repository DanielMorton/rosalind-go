package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input from file
	file, err := os.Open("rosalind_ba10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read emitted string
	scanner.Scan()
	emittedString := scanner.Text()

	// Skip the separator line
	scanner.Scan()

	// Read alphabet
	scanner.Scan()
	alphabet := strings.Fields(scanner.Text())

	// Skip the separator line
	scanner.Scan()

	// Read hidden path
	scanner.Scan()
	hiddenPath := scanner.Text()

	// Skip the separator line
	scanner.Scan()

	// Read states
	scanner.Scan()
	states := strings.Fields(scanner.Text())

	// Skip the separator line
	scanner.Scan()

	// Read emission matrix
	emissionMatrix := make(map[string]map[string]float64)
	for _, state := range states {
		emissionMatrix[state] = make(map[string]float64)
	}

	scanner.Scan() // Skip the header line
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		state := line[0]
		for i, prob := range line[1:] {
			p, _ := strconv.ParseFloat(prob, 64)
			emissionMatrix[state][alphabet[i]] = p
		}
	}

	// Calculate probability
	probability := 1.0
	for i := 0; i < len(emittedString); i++ {
		state := string(hiddenPath[i])
		symbol := string(emittedString[i])
		probability *= emissionMatrix[state][symbol]
	}

	// Output result
	fmt.Printf("%.14e\n", probability)
}
