package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	n int
	k int
}

func solve(pairs []pair) string {
	res := ""
	for _, p := range pairs {
		if p.n%2 == 0 && p.k%2 != 0 {
			res += "NO\n"
		} else if p.n%2 != 0 && p.k%2 == 0 {
			res += "NO\n"
		} else if sumOdd(p.n, p.k) {
			res += "YES\n"
		} else {
			res += "NO\n"
		}
	}
	return res
}

func sumOdd(n, k int) bool {
	sum := 0
	odd := 1
	for sum < n && k != 1 {
		sum += odd
		odd += 2
		k--
	}
	return (sum < n) && ((n-sum)%2 == 1) && (n-sum > odd-2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
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

	lineCount := readInt(r)
	pairs := make([]pair, lineCount)
	for i := 0; i < lineCount; i++ {
		line := readIntSlice(r)
		pairs[i] = pair{
			n: line[0],
			k: line[1],
		}
	}
	result := solve(pairs)
	w.WriteString(fmt.Sprintf("%v", result))
}
