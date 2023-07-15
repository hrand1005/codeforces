package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(laptops [][]int) string {
	priceVal := make(map[int]int)
	prices := make([]int, len(laptops))
	for i, l := range laptops {
		priceVal[l[0]] = l[1]
		prices[i] = l[0]
	}
	sort.Ints(prices)

	maxVal := priceVal[prices[0]]
	for _, p := range prices[1:] {
		if priceVal[p] < maxVal {
			return "Happy Alex"
		} else {
			maxVal = priceVal[p]
		}
	}
	return "Poor Alex"
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

	n := readInt(r)
	laptops := make([][]int, n)
	for i := 0; i < n; i++ {
		laptops[i] = readIntSlice(r)
	}

	result := solve(laptops)
	w.WriteString(fmt.Sprintf("%v\n", result))
}
