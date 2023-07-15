package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(s string, queries [][]int) []int {
	table := make([]int, len(s)+1)
	if s[0] == s[1] {
		table[1] = 1
	}
	for i := 2; i < len(s); i++ {
		if s[i-1] == s[i] {
			table[i] = table[i-1] + 1
		} else {
			table[i] = table[i-1]
		}
	}
	table[len(s)] = table[len(s)-1]

	responses := make([]int, len(queries))
	for i, q := range queries {
		lo, hi := q[0], q[1]
		responses[i] = table[hi-1] - table[lo-1]
	}
	return responses
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func stringToIntSlice(s string) []int {
	nums := strings.Split(s, " ")
	ints := make([]int, len(nums))
	for i, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalf("failed to parse int: %v\n", err)
		}
		ints[i] = n
	}
	return ints
}

func stringSliceToIntSlice(nums []string) []int {
	ints := make([]int, len(nums))
	for i, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalf("failed to parse int: %v\n", err)
		}
		ints[i] = n
	}
	return ints
}

func readString(r *bufio.Reader) string {
	s, err := r.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read string: %v\n", err)
	}
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\n", "")
	return s
}

func readInt(r *bufio.Reader) int {
	s := readString(r)
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to parse int: %v\n", err)
	}
	return n
}

func readStringSlice(r *bufio.Reader) []string {
	s := readString(r)
	return strings.Split(s, " ")
}

func readIntSlice(r *bufio.Reader) []int {
	nums := readStringSlice(r)
	return stringSliceToIntSlice(nums)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	s := readString(r)
	n := readInt(r)
	queries := make([][]int, n)
	for i := 0; i < n; i++ {
		queries[i] = readIntSlice(r)
	}

	result := solve(s, queries)
	for _, r := range result {
		w.WriteString(fmt.Sprintf("%v\n", r))
	}
}
