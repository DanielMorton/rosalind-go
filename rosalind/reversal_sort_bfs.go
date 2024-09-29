package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Permutation []int

type Node struct {
	perm     Permutation
	path     [][2]int
	distance int
}

func main() {
	// Read input from file
	file, err := os.Open("rosalind_sort.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var permutations []Permutation

	for scanner.Scan() {
		line := scanner.Text()
		perm := stringToPermutation(line)
		permutations = append(permutations, perm)
	}

	if len(permutations) != 2 {
		fmt.Println("Invalid input: expected two permutations")
		return
	}

	pi, gamma := permutations[0], permutations[1]
	distance, reversals := bidirectionalBFS(pi, gamma)

	// Output results
	fmt.Println(distance)
	for _, rev := range reversals {
		fmt.Printf("%d %d\n", rev[0], rev[1])
	}
}

func stringToPermutation(s string) Permutation {
	strNums := strings.Fields(s)
	perm := make(Permutation, len(strNums))
	for i, strNum := range strNums {
		num, _ := strconv.Atoi(strNum)
		perm[i] = num
	}
	return perm
}

func bidirectionalBFS(start, end Permutation) (int, [][2]int) {
	forwardQueue := []Node{{perm: start, path: [][2]int{}, distance: 0}}
	backwardQueue := []Node{{perm: end, path: [][2]int{}, distance: 0}}
	forwardVisited := make(map[string]Node)
	backwardVisited := make(map[string]Node)

	for len(forwardQueue) > 0 && len(backwardQueue) > 0 {
		// Forward BFS
		currentForward := forwardQueue[0]
		forwardQueue = forwardQueue[1:]
		forwardKey := permToString(currentForward.perm)

		if node, exists := backwardVisited[forwardKey]; exists {
			return currentForward.distance + node.distance, mergePaths(currentForward.path, node.path)
		}

		if _, visited := forwardVisited[forwardKey]; !visited {
			forwardVisited[forwardKey] = currentForward
			neighbors := generateNeighbors(currentForward)
			forwardQueue = append(forwardQueue, neighbors...)
		}

		// Backward BFS
		currentBackward := backwardQueue[0]
		backwardQueue = backwardQueue[1:]
		backwardKey := permToString(currentBackward.perm)

		if node, exists := forwardVisited[backwardKey]; exists {
			return node.distance + currentBackward.distance, mergePaths(node.path, currentBackward.path)
		}

		if _, visited := backwardVisited[backwardKey]; !visited {
			backwardVisited[backwardKey] = currentBackward
			neighbors := generateNeighbors(currentBackward)
			backwardQueue = append(backwardQueue, neighbors...)
		}
	}

	return -1, nil // No solution found
}

func generateNeighbors(node Node) []Node {
	neighbors := []Node{}
	n := len(node.perm)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			newPerm := make(Permutation, n)
			copy(newPerm, node.perm)
			reverse(newPerm, i, j)

			newPath := make([][2]int, len(node.path))
			copy(newPath, node.path)
			newPath = append(newPath, [2]int{i + 1, j + 1})

			neighbors = append(neighbors, Node{
				perm:     newPerm,
				path:     newPath,
				distance: node.distance + 1,
			})
		}
	}

	return neighbors
}

func reverse(perm Permutation, start, end int) {
	for start < end {
		perm[start], perm[end] = perm[end], perm[start]
		start++
		end--
	}
}

func permToString(perm Permutation) string {
	strs := make([]string, len(perm))
	for i, v := range perm {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, ",")
}

func mergePaths(forward, backward [][2]int) [][2]int {
	result := make([][2]int, len(forward)+len(backward))
	copy(result, forward)
	for i := len(backward) - 1; i >= 0; i-- {
		result[len(forward)+len(backward)-1-i] = backward[i]
	}
	return result
}
