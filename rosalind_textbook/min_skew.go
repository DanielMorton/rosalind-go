package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// findMinimumSkewPositions returns all positions where the skew is minimized
func findMinimumSkewPositions(reader *bufio.Reader) ([]int, error) {
	skew := 0
	minSkew := 0
	var minPositions []int
	position := 0

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		position++
		switch r {
		case 'G':
			skew++
		case 'C':
			skew--
		}

		if skew < minSkew {
			minSkew = skew
			minPositions = []int{position}
		} else if skew == minSkew {
			minPositions = append(minPositions, position)
		}
	}

	return minPositions, nil
}

func main() {
	// Open the input file
	file, err := os.Open("rosalind_ba1f.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)

	// Find positions of minimum skew
	minPositions, err := findMinimumSkewPositions(reader)
	if err != nil {
		fmt.Println("Error processing input:", err)
		return
	}

	// Print the result
	for i, pos := range minPositions {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(pos)
	}
	fmt.Println()
}
