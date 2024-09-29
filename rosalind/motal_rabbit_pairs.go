package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func mortalFibonacci(n, m int) *big.Int {
	// Initialize an array to store the number of rabbit pairs at each age using big.Int
	ages := make([]*big.Int, m)
	for i := range ages {
		ages[i] = big.NewInt(0)
	}
	// Start with one pair of newborn rabbits
	ages[0].SetInt64(1)

	// Simulate month by month
	for month := 1; month < n; month++ {
		// Number of new rabbits is the sum of all rabbits that are at least 1 month old but younger than m months
		newborns := big.NewInt(0)
		for i := 1; i < m; i++ {
			newborns.Add(newborns, ages[i])
		}

		// Shift all age groups to the next age, and the oldest rabbits (ages[m-1]) die
		for i := m - 1; i > 0; i-- {
			ages[i].Set(ages[i-1])
		}

		// Add the newborns to the 0-month-old group
		ages[0].Set(newborns)
	}

	// Sum up all the remaining rabbits
	totalRabbits := big.NewInt(0)
	for i := 0; i < m; i++ {
		totalRabbits.Add(totalRabbits, ages[i])
	}

	return totalRabbits
}

func main() {
	// Read input from file
	data, err := os.ReadFile("rosalind_fibd.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Parse the input data
	input := strings.TrimSpace(string(data))
	values := strings.Split(input, " ")
	n, _ := strconv.Atoi(values[0])
	m, _ := strconv.Atoi(values[1])

	// Call the function to calculate the result
	result := mortalFibonacci(n, m)

	// Print the result
	fmt.Println(result.String())
}
