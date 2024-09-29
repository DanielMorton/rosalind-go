package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func calculateNXX(lengths []int, totalLength int, percentage float64) int {
	sort.Sort(sort.Reverse(sort.IntSlice(lengths)))
	threshold := int(float64(totalLength) * percentage)
	sum := 0
	for _, length := range lengths {
		sum += length
		if sum >= threshold {
			return length
		}
	}
	return 0
}

func main() {
	file, err := os.Open("rosalind_asmq.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lengths []int
	totalLength := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contig := scanner.Text()
		length := len(contig)
		lengths = append(lengths, length)
		totalLength += length
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	n50 := calculateNXX(lengths, totalLength, 0.5)
	n75 := calculateNXX(lengths, totalLength, 0.75)

	fmt.Printf("%d %d\n", n50, n75)
}
