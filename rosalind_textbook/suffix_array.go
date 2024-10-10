package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type suffixArray struct {
	text string
	sa   []int
}

func newSuffixArray(text string) *suffixArray {
	sa := &suffixArray{
		text: text,
		sa:   make([]int, len(text)),
	}
	for i := range sa.sa {
		sa.sa[i] = i
	}
	sort.Slice(sa.sa, func(i, j int) bool {
		return sa.less(sa.sa[i], sa.sa[j])
	})
	return sa
}

func (sa *suffixArray) less(i, j int) bool {
	for k := 0; k < len(sa.text)-i && k < len(sa.text)-j; k++ {
		if sa.text[i+k] != sa.text[j+k] {
			return sa.text[i+k] < sa.text[j+k]
		}
	}
	return i > j
}

func constructSuffixArray(text string) []int {
	sa := newSuffixArray(text)
	return sa.sa
}

func main() {
	// Read input from file
	content, err := ioutil.ReadFile("rosalind_ba9g.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Trim whitespace and newlines
	text := strings.TrimSpace(string(content))

	// Construct the suffix array
	suffixArray := constructSuffixArray(text)

	// Print the result
	for i, v := range suffixArray {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
