package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

type suffixArray struct {
	suffixes []int
	s        string
}

func (sa suffixArray) Len() int           { return len(sa.suffixes) }
func (sa suffixArray) Less(i, j int) bool { return sa.s[sa.suffixes[i]:] < sa.s[sa.suffixes[j]:] }
func (sa suffixArray) Swap(i, j int)      { sa.suffixes[i], sa.suffixes[j] = sa.suffixes[j], sa.suffixes[i] }

func buildSuffixArray(s string) []int {
	sa := &suffixArray{
		suffixes: make([]int, len(s)),
		s:        s,
	}
	for i := range sa.suffixes {
		sa.suffixes[i] = i
	}
	sort.Sort(sa)
	return sa.suffixes
}

func buildLCPArray(s string, sa []int) []int {
	n := len(s)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		rank[sa[i]] = i
	}

	lcp := make([]int, n-1)
	k := 0
	for i := 0; i < n; i++ {
		if rank[i] == n-1 {
			k = 0
			continue
		}
		j := sa[rank[i]+1]
		for i+k < n && j+k < n && s[i+k] == s[j+k] {
			k++
		}
		lcp[rank[i]] = k
		if k > 0 {
			k--
		}
	}
	return lcp
}

func countSubstrings(s string) int {
	n := len(s)
	sa := buildSuffixArray(s)
	lcp := buildLCPArray(s, sa)

	total := n * (n + 1) / 2
	for _, v := range lcp {
		total -= v
	}
	return total
}

func maxSubstringCount(a, n int) int {
	total := 0
	for k := 1; k <= n; k++ {
		total += int(math.Min(float64(math.Pow(float64(a), float64(k))), float64(n-k+1)))
	}
	return total
}

func linguisticComplexity(s string) float64 {
	n := len(s)
	a := 4 // DNA alphabet size
	sub_s := countSubstrings(s)
	m_a_n := maxSubstringCount(a, n)
	return float64(sub_s) / float64(m_a_n)
}

func main() {
	file, err := os.Open("rosalind_ling.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 1024*1024) // 1MB buffer

	var s string
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		s += string(line)
		if !isPrefix {
			break
		}
	}

	lc := linguisticComplexity(s)
	fmt.Printf("%.3f\n", lc)
}
